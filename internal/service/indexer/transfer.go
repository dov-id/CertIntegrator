package indexer

import (
	"math/big"

	"github.com/dov-id/cert-integrator-svc/contracts"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iden3/go-merkletree-sql/v2"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (i *Indexer) handleIssuerTransferLog(eventLog types.Log, client *ethclient.Client) error {
	i.log.WithField("address", eventLog.Address.Hex()).Debugf("start handling transfer event")

	issuer, err := contracts.NewIssuer(eventLog.Address, client)
	if err != nil {
		return errors.Wrap(err, "failed to create new issuer instance")
	}

	event, err := issuer.ParseTransfer(eventLog)
	if err != nil {
		return errors.Wrap(err, "failed to parse transfer event data")
	}

	contract, err := i.ContractsQ.FilterByAddresses(eventLog.Address.Hex()).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get contract from database")
	}

	if contract == nil {
		contract, err = i.ContractsQ.Insert(data.Contract{
			Name:    IssuerContract,
			Address: eventLog.Address.Hex(),
		})
		if err != nil {
			return errors.Wrap(err, "failed to save new contract")
		}
	}

	treeStorage := postgres.NewPGDBStorage(i.cfg.DB(), contract.Id)

	mTree, err := merkletree.NewMerkleTree(i.ctx, treeStorage, 100)
	if err != nil {
		return errors.Wrap(err, "failed to create merkle tree")
	}

	if event.From.Hex() == ZeroAddress {
		err = i.handleMint(mTree, event, int64(eventLog.BlockNumber))
		if err != nil {
			return errors.Wrap(err, "failed to handle mint event")
		}
		return nil
	}

	err = i.handleTransfer(mTree, event, int64(eventLog.BlockNumber), treeStorage)
	if err != nil {
		return errors.Wrap(err, "failed to handle transfer event")
	}
	return nil
}

func (i *Indexer) handleTransfer(mTree *merkletree.MerkleTree, event *contracts.IssuerTransfer, blockNumber int64, treeStorage *postgres.Storage) error {
	receiver := event.To.Big()

	_, leafValue, _, err := mTree.Get(i.ctx, receiver)
	if err != nil {
		return errors.Wrap(err, "failed to get leaf")
	}

	value := leafValue.Int64() - 1
	if value < 1 {
		err = i.completelyDeleteKey(mTree, event, treeStorage)
		if err != nil {
			return err
		}
	} else {
		_, err = mTree.Update(i.ctx, receiver, big.NewInt(value))
		if err != nil {
			return errors.Wrap(err, "failed to update leaf in merkle tree")
		}
	}

	err = i.ContractsQ.FilterByAddresses(event.Raw.Address.Hex()).Update(data.ContractToUpdate{
		Block: &blockNumber,
	})
	if err != nil {
		return errors.Wrap(err, "failed to save last handled block")
	}

	i.log.WithField("address", event.Raw.Address.Hex()).Debugf("finish handling transfer event")
	return nil
}

func (i *Indexer) handleMint(mTree *merkletree.MerkleTree, event *contracts.IssuerTransfer, blockNumber int64) error {
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
	} else {
		value := leafValue.Int64() + 1
		_, err = mTree.Update(i.ctx, receiver, big.NewInt(value))
		if err != nil {
			return errors.Wrap(err, "failed to update leaf in merkle tree")
		}
	}

	err = i.ContractsQ.FilterByAddresses(event.Raw.Address.Hex()).Update(data.ContractToUpdate{
		Block: &blockNumber,
	})
	if err != nil {
		return errors.Wrap(err, "failed to update last handled block")
	}

	i.log.WithField("address", event.Raw.Address.Hex()).Debugf("finish handling mint event")
	return nil
}

func (i *Indexer) completelyDeleteKey(mTree *merkletree.MerkleTree, event *contracts.IssuerTransfer, treeStorage *postgres.Storage) error {
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
		return errors.New("no contract was found")
	}

	treeStorage = postgres.NewPGDBStorage(i.cfg.DB(), contract.Id)
	newMTree, err := merkletree.NewMerkleTree(i.ctx, treeStorage, 100)
	if err != nil {
		return errors.Wrap(err, "failed to create merkle tree")
	}

	err = newMTree.ImportDumpedLeafs(i.ctx, dump)
	if err != nil {
		return errors.Wrap(err, "failed to import dumped leafs")
	}

	return nil
}
