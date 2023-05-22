package indexer

import (
	"context"

	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"github.com/dov-id/cert-integrator-svc/internal/helpers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iden3/go-merkletree-sql/v2"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

// TODO: replace with generated code
type LogDeploy struct {
	//event Deploy(string, address indexed contract);
	CourseName    string
	CourseAddress common.Address
}

func (i *Indexer) handleFabricDeployLog(eventLog types.Log, _ *ethclient.Client) error {
	i.log.Debugf("start handling deploy event")

	var event LogDeploy

	event.CourseName = eventLog.Topics[1].String()
	event.CourseAddress = common.HexToAddress(eventLog.Topics[2].Hex())

	newContract, err := i.ContractsQ.Insert(data.Contract{
		Name:    event.CourseName,
		Address: event.CourseAddress.Hex(),
	})
	if err != nil {
		return errors.Wrap(err, "failed to save new contract")
	}

	treeStorage := postgres.NewPGDBStorage(i.cfg.DB(), newContract.Id)

	_, err = merkletree.NewMerkleTree(i.ctx, treeStorage, 100)
	if err != nil {
		return errors.Wrap(err, "failed to create new merkle tree")
	}

	blockNumber := int64(eventLog.BlockNumber)
	err = i.recreateIssuerRunner(blockNumber, newContract.Address)
	if err != nil {
		return errors.Wrap(err, "failed to recreate issuers runner")
	}

	err = i.ContractsQ.FilterByAddresses(eventLog.Address.Hex()).Update(data.ContractToUpdate{
		Block: &blockNumber,
	})
	if err != nil {
		return errors.Wrap(err, "failed to save last handled block")
	}

	i.log.Debugf("finish handling deploy event")
	return nil
}

func (i *Indexer) recreateIssuerRunner(block int64, address string) error {
	if i.Cancel == nil {
		return errors.New("ctx cancel function in nil")
	}

	i.Cancel()

	contracts, err := i.ContractsQ.Select()
	if err != nil {
		return errors.Wrap(err, "failed to select contracts")
	}

	blocks, addresses := helpers.SeparateDataContractArrays(contracts)
	cancelCtx, cancelFn := context.WithCancel(i.ctx)

	NewIndexer(
		i.cfg,
		cancelCtx,
		append(addresses, address),
		append(blocks, block),
		nil,
	).Run(i.ctx)

	i.Cancel = cancelFn
	return nil
}
