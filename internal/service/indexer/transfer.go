package indexer

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"sync"

	"github.com/dov-id/cert-integrator-svc/contracts"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iden3/go-merkletree-sql/v2"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (i *indexer) handleIssuerTransferLog(eventLog types.Log, client *ethclient.Client) error {
	i.log.WithField("address", eventLog.Address.Hex()).Debugf("start handling transfer event")

	issuer, err := contracts.NewTokenContract(eventLog.Address, client)
	if err != nil {
		return errors.Wrap(err, "failed to create new issuer instance")
	}

	event, err := issuer.ParseTransfer(eventLog)
	if err != nil {
		return errors.Wrap(err, "failed to parse transfer event data")
	}

	contract, err := i.ContractsQ.FilterByAddresses(event.Raw.Address.Hex()).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get contract from database")
	}

	if contract == nil {
		contract, err = i.ContractsQ.Insert(data.Contract{
			Name:    IssuerContract,
			Address: event.Raw.Address.Hex(),
			Type:    data.ISSUER,
		})
		if err != nil {
			return errors.Wrap(err, "failed to save new contract")
		}
	}

	treeStorage := postgres.NewStorage(i.cfg.DB(), contract.Id)

	mTree, err := merkletree.NewMerkleTree(i.ctx, treeStorage, data.MaxMTreeLevel)
	if err != nil {
		return errors.Wrap(err, "failed to create merkle tree")
	}

	if event.From.Hex() == ZeroAddress {
		err = i.handleMint(mTree, event, int64(event.Raw.BlockNumber))
		if err != nil {
			return errors.Wrap(err, "failed to handle mint event")
		}
		return nil
	}

	err = i.handleTransfer(mTree, event, int64(event.Raw.BlockNumber), treeStorage)
	if err != nil {
		return errors.Wrap(err, "failed to handle transfer event")
	}

	return nil
}

func (i *indexer) handleTransfer(mTree *merkletree.MerkleTree, event *contracts.TokenContractTransfer, blockNumber int64, treeStorage *postgres.Storage) error {
	receiver := event.To.Big()

	_, leafValue, _, err := mTree.Get(i.ctx, receiver)
	if err != nil {
		return errors.Wrap(err, "failed to get leaf")
	}

	value := leafValue.Int64() - 1
	if value < 1 {
		err = i.completelyDeleteKey(mTree, event, treeStorage)
		if err != nil {
			return errors.Wrap(err, "failed to fully delete key")
		}

		return i.updateContractsStates(event, blockNumber, mTree.Root(), "finish handling transfer event")
	}
	_, err = mTree.Update(i.ctx, receiver, big.NewInt(value))
	if err != nil {
		return errors.Wrap(err, "failed to update leaf in merkle tree")
	}

	return i.updateContractsStates(event, blockNumber, mTree.Root(), "finish handling transfer event")
}

func (i *indexer) updateContractsStates(event *contracts.TokenContractTransfer, blockNumber int64, root *merkletree.Hash, msg string) error {
	err := i.ContractsQ.FilterByAddresses(event.Raw.Address.Hex()).Update(data.ContractToUpdate{
		Block: &blockNumber,
	})
	if err != nil {
		return errors.Wrap(err, "failed to update last handled block")
	}

	err = i.publish(event.Raw.Address.Hex(), root)
	if err != nil {
		return errors.Wrap(err, "failed to publish")
	}

	i.log.WithField("address", event.Raw.Address.Hex()).Debugf(msg)
	return nil
}

func (i *indexer) handleMint(mTree *merkletree.MerkleTree, event *contracts.TokenContractTransfer, blockNumber int64) error {
	receiver := event.To.Big()

	_, leafValue, _, err := mTree.Get(i.ctx, receiver)
	if err != nil && err != merkletree.ErrKeyNotFound {
		return errors.Wrap(err, "failed to get leaf")
	}

	if err == merkletree.ErrKeyNotFound {
		err = mTree.Add(i.ctx, receiver, big.NewInt(1))
		if err != nil {
			return errors.Wrap(err, "failed to add new leaf in merkle tree")
		}

		return i.updateContractsStates(event, blockNumber, mTree.Root(), "finish handling mint event")
	}

	value := leafValue.Int64() + 1
	_, err = mTree.Update(i.ctx, receiver, big.NewInt(value))
	if err != nil {
		return errors.Wrap(err, "failed to update leaf in merkle tree")
	}

	return i.updateContractsStates(event, blockNumber, mTree.Root(), "finish handling mint event")
}

func (i *indexer) completelyDeleteKey(mTree *merkletree.MerkleTree, event *contracts.TokenContractTransfer, treeStorage *postgres.Storage) error {
	err := mTree.Delete(i.ctx, event.To.Big())
	if err != nil {
		return errors.Wrap(err, "failed to delete address from merkle tree")
	}

	dump, err := mTree.DumpLeafs(i.ctx, mTree.Root())
	if err != nil {
		return errors.Wrap(err, "failed to make dump of merkle tree")
	}

	err = treeStorage.DeleteMTree(i.ctx)
	if err != nil {
		return errors.Wrap(err, "failed to delete mtree")
	}

	contract, err := i.ContractsQ.FilterByAddresses(event.Raw.Address.Hex()).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get contract")
	}
	if contract == nil {
		return errors.New(data.NoContractErr)
	}

	treeStorage = postgres.NewStorage(i.cfg.DB(), contract.Id)
	newMTree, err := merkletree.NewMerkleTree(i.ctx, treeStorage, data.MaxMTreeLevel)
	if err != nil {
		return errors.Wrap(err, "failed to create merkle tree")
	}

	err = newMTree.ImportDumpedLeafs(i.ctx, dump)
	if err != nil {
		return errors.Wrap(err, "failed to import dumped leafs")
	}

	return nil
}

func (i *indexer) publish(name string, root *merkletree.Hash) error {
	err := i.sendUpdates(i.EthereumClient, name, root, i.CertIntegratorEthereum)
	if err != nil {
		return errors.Wrap(err, "failed to publish in ethereum")
	}

	err = i.sendUpdates(i.PolygonClient, name, root, i.CertIntegratorPolygon)
	if err != nil {
		return errors.Wrap(err, "failed to publish in polygon")
	}

	err = i.sendUpdates(i.QClient, name, root, i.CertIntegratorQ)
	if err != nil {
		return errors.Wrap(err, "failed to publish in q")
	}

	return nil
}

func (i *indexer) sendUpdates(client *ethclient.Client, name string, root *merkletree.Hash, certIntegrator *contracts.CertIntegratorContract) error {
	auth, err := i.getAuth(client)
	if err != nil {
		return errors.Wrap(err, "failed to get auth options")
	}

	var course = []byte(name)
	var state = []byte(hexutil.Encode(common.LeftPadBytes(
		root.BigInt().Bytes(),
		32,
	)))

	err = sendUpdateCourseState(client, certIntegrator, auth, course, state)
	if err != nil {
		return errors.Wrap(err, "failed to update course state")
	}

	return nil
}

func sendUpdateCourseState(client *ethclient.Client, certIntegrator *contracts.CertIntegratorContract, auth *bind.TransactOpts, course []byte, state []byte) error {
	transaction, err := certIntegrator.UpdateCourseState(auth, [][]byte{course}, [][]byte{state})
	if err != nil {
		if err.Error() == data.ReplacementTxUnderpricedErr {
			auth.Nonce = big.NewInt(auth.Nonce.Int64() + 1)
			return sendUpdateCourseState(client, certIntegrator, auth, course, state)
		}

		return errors.Wrap(err, "failed to update course state")
	}

	err = waitForTransactionMined(client, transaction)
	if err != nil {
		return errors.Wrap(err, "failed to wait for tx mined")
	}

	return nil
}

func waitForTransactionMined(client *ethclient.Client, transaction *types.Transaction) error {
	var (
		err   error
		mined = make(chan struct{})
		ctx   = context.Background()
	)

	go func() {
		_, err = bind.WaitMined(ctx, client, transaction)
		close(mined)
	}()

	select {
	case <-mined:
		if err != nil {
			return errors.Wrap(err, "failed to mine transaction")
		}
	}

	return nil
}

func (i *indexer) getAuth(client *ethclient.Client) (*bind.TransactOpts, error) {
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain id")
	}

	privateKey, fromAddress, err := i.getKeys()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get keys")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transaction signer")
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get nonce")
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to suggest gas price")
	}

	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	auth.Nonce = big.NewInt(int64(nonce))

	return auth, nil
}

func (i *indexer) getKeys() (*ecdsa.PrivateKey, common.Address, error) {
	var once sync.Once
	var privateKey *ecdsa.PrivateKey
	var fromAddress common.Address
	var err error

	once.Do(func() {
		privateKey, err = crypto.HexToECDSA(i.cfg.Networks().Networks[data.MetamaskNetwork].PrivateKey)
		if err != nil {
			err = errors.Wrap(err, "failed to convert hex to ecdsa")
			return
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			err = errors.New(data.FailedToCastKeyErr)
			return
		}

		fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
		return
	})

	return privateKey, fromAddress, nil
}
