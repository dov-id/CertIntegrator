package indexer

import (
	"context"

	"github.com/dov-id/cert-integrator-svc/contracts"
	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"github.com/dov-id/cert-integrator-svc/internal/helpers"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Indexer interface {
	Run(ctx context.Context)
}

type indexer struct {
	cfg config.Config
	ctx context.Context
	log *logan.Entry

	Addresses  []string
	Blocks     []int64
	ContractsQ data.Contracts

	Cancel context.CancelFunc

	EthereumClient *ethclient.Client
	PolygonClient  *ethclient.Client
	QClient        *ethclient.Client

	CertIntegratorEthereum *contracts.CertIntegratorContract
	CertIntegratorPolygon  *contracts.CertIntegratorContract
	CertIntegratorQ        *contracts.CertIntegratorContract
}

func Run(cfg config.Config, ctx context.Context) {
	cancelCtx, cancelFn := context.WithCancel(ctx)
	err := updateContractsDB(postgres.NewContractsQ(cfg.DB()), cfg.CertificatesIssuer().List, IssuerContract, data.ISSUER)
	if err != nil {
		panic(err)
	}
	blocks, addresses := helpers.SeparateContractArrays(cfg.CertificatesIssuer().List)
	NewIndexer(
		cfg,
		cancelCtx,
		addresses,
		blocks,
		nil,
	).Run(ctx)

	err = updateContractsDB(postgres.NewContractsQ(cfg.DB()), cfg.CertificatesFabric().List, FabricContract, data.FABRIC)
	if err != nil {
		panic(err)
	}
	blocks, addresses = helpers.SeparateContractArrays(cfg.CertificatesFabric().List)
	NewIndexer(
		cfg,
		ctx,
		addresses,
		blocks,
		cancelFn,
	).Run(ctx)
}

func updateContractsDB(contractsQ data.Contracts, list []config.Contract, name string, types data.ContractType) error {
	for i := range list {
		contract, err := contractsQ.FilterByAddresses(list[i].Address).Get()
		if err != nil {
			return errors.Wrap(err, "failed to get contract from database")
		}

		if contract != nil {
			continue
		}

		contract, err = contractsQ.Insert(data.Contract{
			Name:    name,
			Address: list[i].Address,
			Block:   list[i].FromBlock,
			Type:    types,
		})
		if err != nil {
			return errors.Wrap(err, "failed to save new contract")
		}
	}
	return nil
}
