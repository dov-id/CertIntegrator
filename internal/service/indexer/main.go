package indexer

import (
	"context"
	"sync"

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
	wg     *sync.WaitGroup

	EthereumClient *ethclient.Client
	PolygonClient  *ethclient.Client
	QClient        *ethclient.Client

	CertIntegratorEthereum *contracts.CertIntegratorContract
	CertIntegratorPolygon  *contracts.CertIntegratorContract
	CertIntegratorQ        *contracts.CertIntegratorContract
}

func Run(cfg config.Config, ctx context.Context) {
	cancelCtx, cancelFn := context.WithCancel(ctx)
	var wg sync.WaitGroup

	blocks, addresses, err := updAndGetContractsInfo(postgres.NewContractsQ(cfg.DB()), cfg.CertificatesIssuer().List, IssuerContract, data.ISSUER)
	if err != nil {
		panic(errors.Wrap(err, "failed to update and retrieve contracts info"))
	}

	wg.Add(1)

	NewIndexer(
		cfg,
		cancelCtx,
		addresses,
		blocks,
		nil,
		&wg,
	).Run(cancelCtx)

	blocks, addresses, err = updAndGetContractsInfo(postgres.NewContractsQ(cfg.DB()), cfg.CertificatesFabric().List, FabricContract, data.FABRIC)
	if err != nil {
		panic(errors.Wrap(err, "failed to update and retrieve contracts info"))
	}

	NewIndexer(
		cfg,
		ctx,
		addresses,
		blocks,
		cancelFn,
		&wg,
	).Run(ctx)
}

func updAndGetContractsInfo(contractsQ data.Contracts, list []config.Contract, name string, types data.ContractType) ([]int64, []string, error) {
	err := updateContractsDB(contractsQ, list, name, types)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to update contracts in db")
	}

	cfgBlocks, cfgAddresses := helpers.SeparateContractArrays(list)

	dbContracts, err := contractsQ.FilterByTypes(types).Select()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get contract from database")
	}

	dbBlocks, dbAddresses := helpers.SeparateDataContractArrays(dbContracts)

	return helpers.RemoveDuplicatesInt64Arr(append(cfgBlocks, dbBlocks...)),
		helpers.RemoveDuplicatesStringsArr(append(cfgAddresses, dbAddresses...)),
		nil
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
