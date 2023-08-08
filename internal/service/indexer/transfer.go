package indexer

import (
	"context"
	"math/big"

	"github.com/dov-id/cert-integrator-svc/contracts"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iden3/go-merkletree-sql/v2"
	pkgErrors "github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (i *indexer) handleIssuerTransferLog(ctx context.Context, eventLog types.Log, client *ethclient.Client) error {
	i.log.WithField("address", eventLog.Address.Hex()).Debugf("start handling transfer event")

	issuer, err := contracts.NewTokenContract(eventLog.Address, client)
	if err != nil {
		return errors.Wrap(err, "failed to create new issuer instance")
	}

	event, err := issuer.ParseTransfer(eventLog)
	if err != nil {
		return errors.Wrap(err, "failed to parse transfer event data")
	}

	contract, err := i.MasterQ.ContractsQ().FilterByAddresses(event.Raw.Address.Hex()).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get contract from database")
	}

	if contract == nil {
		contract, err = i.MasterQ.ContractsQ().Insert(data.Contract{
			Name:    IssuerContract,
			Address: event.Raw.Address.Hex(),
			Type:    data.Issuer,
		})
		if err != nil {
			return errors.Wrap(err, "failed to save new contract")
		}
	}

	treeStorage := postgres.NewStorage(i.cfg.DB(), contract.Id)

	mTree, err := merkletree.NewMerkleTree(ctx, treeStorage, data.MaxMTreeLevel)
	if err != nil {
		return errors.Wrap(err, "failed to create merkle tree")
	}

	err = i.MasterQ.UsersQ().Insert(data.User{
		Address:    event.To.Hex(),
		ContractId: contract.Id,
	})
	if err != nil {
		return errors.Wrap(err, "failed to insert user")
	}

	if event.From.Hex() == ZeroAddress {
		err = i.handleMint(ctx, mTree, event, int64(event.Raw.BlockNumber))
		if err != nil {
			return errors.Wrap(err, "failed to handle mint event")
		}
		return nil
	}

	err = i.handleTransfer(ctx, mTree, event, int64(event.Raw.BlockNumber), treeStorage)
	if err != nil {
		return errors.Wrap(err, "failed to handle transfer event")
	}

	return nil
}

func (i *indexer) handleTransfer(
	ctx context.Context,
	mTree *merkletree.MerkleTree,
	event *contracts.TokenContractTransfer,
	blockNumber int64,
	treeStorage *postgres.Storage,
) error {
	receiver := event.To.Big()

	_, leafValue, _, err := mTree.Get(ctx, receiver)
	if err != nil {
		return errors.Wrap(err, "failed to get leaf")
	}

	value := leafValue.Int64() - 1
	if value < 1 {
		err = i.completelyDeleteKey(ctx, mTree, event, treeStorage)
		if err != nil {
			return errors.Wrap(err, "failed to fully delete key")
		}

		return i.updateContractsStates(event, blockNumber, mTree.Root(), "finish handling transfer event")
	}
	_, err = mTree.Update(ctx, receiver, big.NewInt(value))
	if err != nil {
		return errors.Wrap(err, "failed to update leaf in merkle tree")
	}

	return i.updateContractsStates(event, blockNumber, mTree.Root(), "finish handling transfer event")
}

func (i *indexer) handleMint(
	ctx context.Context,
	mTree *merkletree.MerkleTree,
	event *contracts.TokenContractTransfer,
	blockNumber int64,
) error {
	receiver := event.To.Big()

	_, leafValue, _, err := mTree.Get(ctx, receiver)
	if err != nil && !pkgErrors.Is(err, merkletree.ErrKeyNotFound) {
		return errors.Wrap(err, "failed to get leaf")
	}

	if pkgErrors.Is(err, merkletree.ErrKeyNotFound) {
		err = mTree.Add(ctx, receiver, big.NewInt(1))
		if err != nil {
			return errors.Wrap(err, "failed to add new leaf in merkle tree")
		}

		return i.updateContractsStates(event, blockNumber, mTree.Root(), "finish handling mint event")
	}

	value := leafValue.Int64() + 1
	_, err = mTree.Update(ctx, receiver, big.NewInt(value))
	if err != nil {
		return errors.Wrap(err, "failed to update leaf in merkle tree")
	}

	return i.updateContractsStates(event, blockNumber, mTree.Root(), "finish handling mint event")
}

func (i *indexer) completelyDeleteKey(
	ctx context.Context,
	mTree *merkletree.MerkleTree,
	event *contracts.TokenContractTransfer,
	treeStorage *postgres.Storage,
) error {
	err := mTree.Delete(ctx, event.To.Big())
	if err != nil {
		return errors.Wrap(err, "failed to delete address from merkle tree")
	}

	dump, err := mTree.DumpLeafs(ctx, mTree.Root())
	if err != nil {
		return errors.Wrap(err, "failed to make dump of merkle tree")
	}

	err = treeStorage.DeleteMTree(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to delete mtree")
	}

	contract, err := i.MasterQ.ContractsQ().FilterByAddresses(event.Raw.Address.Hex()).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get contract")
	}
	if contract == nil {
		return data.ErrNoContract
	}

	treeStorage = postgres.NewStorage(i.cfg.DB(), contract.Id)
	newMTree, err := merkletree.NewMerkleTree(ctx, treeStorage, data.MaxMTreeLevel)
	if err != nil {
		return errors.Wrap(err, "failed to create merkle tree")
	}

	err = newMTree.ImportDumpedLeafs(ctx, dump)
	if err != nil {
		return errors.Wrap(err, "failed to import dumped leafs")
	}

	return nil
}

func (i *indexer) updateContractsStates(
	event *contracts.TokenContractTransfer,
	blockNumber int64,
	root *merkletree.Hash,
	msg string,
) error {
	err := i.MasterQ.ContractsQ().FilterByAddresses(event.Raw.Address.Hex()).Update(data.ContractToUpdate{
		Block: &blockNumber,
	})
	if err != nil {
		return errors.Wrap(err, "failed to update last handled block")
	}
	i.Blocks[event.Raw.Address.Hex()] = blockNumber

	err = i.MasterQ.TransactionsQ().Insert(data.Transaction{
		Status: data.TransactionStatusPending,
		Course: event.Raw.Address.Hex(),
		State:  root[:],
	})
	if err != nil {
		return errors.Wrap(err, "failed to insert transaction")
	}

	i.log.WithField("address", event.Raw.Address.Hex()).Debugf(msg)
	return nil
}
