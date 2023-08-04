package indexer

import (
	"context"

	"github.com/dov-id/cert-integrator-svc/contracts"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iden3/go-merkletree-sql/v2"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (i *indexer) handleFabricDeployLog(ctx context.Context, eventLog types.Log, client *ethclient.Client) error {
	i.log.WithField("address", eventLog.Address.Hex()).Debugf("start handling deploy event")

	fabric, err := contracts.NewTokenFactoryContract(eventLog.Address, client)
	if err != nil {
		return errors.Wrap(err, "failed to create new issuer instance")
	}

	event, err := fabric.ParseTokenContractDeployed(eventLog)
	if err != nil {
		return errors.Wrap(err, "failed to parse transfer event data")
	}

	contract, err := i.MasterQ.ContractsQ().FilterByAddresses(event.NewTokenContractAddr.Hex()).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get contract")
	}

	if contract != nil {
		i.log.WithField("address", contract.Address).Debugf("contract already exists")

		blockNumber := int64(event.Raw.BlockNumber)
		err = i.MasterQ.ContractsQ().FilterByAddresses(event.Raw.Address.Hex()).Update(data.ContractToUpdate{
			Block: &blockNumber,
		})
		if err != nil {
			return errors.Wrap(err, "failed to update fabric contract")
		}
		i.Blocks[event.Raw.Address.Hex()] = blockNumber

		return nil
	}

	err = i.processNewContract(ctx, event)
	if err != nil {
		return errors.Wrap(err, "failed to process new contract")
	}

	i.log.WithField("address", event.Raw.Address.Hex()).Debugf("finish handling deploy event")
	return nil
}

func (i *indexer) processNewContract(ctx context.Context, event *contracts.TokenFactoryContractTokenContractDeployed) error {
	blockNumber := int64(event.Raw.BlockNumber)

	newContract, err := i.MasterQ.ContractsQ().Insert(data.Contract{
		Name:    event.TokenContractParams.TokenName,
		Address: event.NewTokenContractAddr.Hex(),
		Block:   blockNumber,
		Type:    data.Issuer,
	})
	if err != nil {
		return errors.Wrap(err, "failed to save new contract")
	}
	i.Blocks[event.NewTokenContractAddr.Hex()] = blockNumber

	treeStorage := postgres.NewStorage(i.cfg.DB(), newContract.Id)

	_, err = merkletree.NewMerkleTree(ctx, treeStorage, data.MaxMTreeLevel)
	if err != nil {
		return errors.Wrap(err, "failed to create new merkle tree")
	}

	i.issuerCh <- event.NewTokenContractAddr.Hex()

	err = i.MasterQ.ContractsQ().FilterByAddresses(event.Raw.Address.Hex()).Update(data.ContractToUpdate{
		Block: &blockNumber,
	})
	if err != nil {
		return errors.Wrap(err, "failed to update fabric contract")
	}
	i.Blocks[event.Raw.Address.Hex()] = blockNumber

	return nil
}
