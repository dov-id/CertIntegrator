package indexer

import (
	"context"

	"github.com/dov-id/cert-integrator-svc/contracts"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"github.com/dov-id/cert-integrator-svc/internal/helpers"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iden3/go-merkletree-sql/v2"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (i *Indexer) handleFabricDeployLog(eventLog types.Log, client *ethclient.Client) error {
	i.log.WithField("address", eventLog.Address.Hex()).Debugf("start handling deploy event")

	fabric, err := contracts.NewTokenFactory(eventLog.Address, client)
	if err != nil {
		return errors.Wrap(err, "failed to create new issuer instance")
	}

	event, err := fabric.ParseTokenContractDeployed(eventLog)
	if err != nil {
		return errors.Wrap(err, "failed to parse transfer event data")
	}

	contract, err := i.ContractsQ.FilterByAddresses(event.NewTokenContractAddr.Hex()).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get contract")
	}

	if contract != nil {
		i.log.WithField("address", contract.Address).Debugf("contract already exists")

		blockNumber := int64(event.Raw.BlockNumber)
		err = i.ContractsQ.FilterByAddresses(contract.Address).Update(data.ContractToUpdate{
			Name:  &event.TokenContractParams.TokenName,
			Block: &blockNumber,
		})
		return err
	}

	newContract, err := i.ContractsQ.Insert(data.Contract{
		Name:    event.TokenContractParams.TokenName,
		Address: event.NewTokenContractAddr.Hex(),
		Type:    data.ISSUER,
	})
	if err != nil {
		return errors.Wrap(err, "failed to save new contract")
	}

	treeStorage := postgres.NewPGDBStorage(i.cfg.DB(), newContract.Id)

	_, err = merkletree.NewMerkleTree(i.ctx, treeStorage, 100)
	if err != nil {
		return errors.Wrap(err, "failed to create new merkle tree")
	}

	blockNumber := int64(event.Raw.BlockNumber)
	err = i.recreateIssuerRunner(blockNumber)
	if err != nil {
		return errors.Wrap(err, "failed to recreate issuers runner")
	}

	err = i.ContractsQ.FilterByAddresses(newContract.Address).Update(data.ContractToUpdate{
		Block: &blockNumber,
	})
	if err != nil {
		return errors.Wrap(err, "failed to save last handled block")
	}

	i.log.WithField("address", event.Raw.Address.Hex()).Debugf("finish handling deploy event")
	return nil
}

func (i *Indexer) recreateIssuerRunner(block int64) error {
	if i.Cancel == nil {
		return errors.New("ctx cancel function is nil")
	}

	i.Cancel()

	dbContracts, err := i.ContractsQ.FilterByTypes(data.ISSUER).Select()
	if err != nil {
		return errors.Wrap(err, "failed to select contracts")
	}

	blocks, addresses := helpers.SeparateDataContractArrays(dbContracts)
	cancelCtx, cancelFn := context.WithCancel(i.ctx)

	NewIndexer(
		i.cfg,
		cancelCtx,
		addresses,
		append(blocks, block),
		nil,
	).Run(i.ctx)

	i.Cancel = cancelFn
	return nil
}
