package indexer

import (
	"context"

	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"github.com/dov-id/cert-integrator-svc/internal/helpers"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func updAndGetContractsInfo(contractsQ data.Contracts, list []config.Contract, types data.ContractType) (map[string]int64, []string, error) {
	err := updateContractsDB(contractsQ, list, types)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to update contracts in db")
	}

	dbContracts, err := contractsQ.FilterByTypes(types).Select()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get contract from database")
	}

	blocks := make(map[string]int64)
	addresses := make([]string, len(dbContracts))

	for i := 0; i < len(dbContracts); i++ {
		blocks[dbContracts[i].Address] = dbContracts[i].Block
		addresses[i] = dbContracts[i].Address
	}

	return blocks, addresses, nil
}

func updateContractsDB(contractsQ data.Contracts, list []config.Contract, types data.ContractType) error {
	for i := range list {
		contract, err := contractsQ.FilterByAddresses(list[i].Address).Get()
		if err != nil {
			return errors.Wrap(err, "failed to get contract from database")
		}

		if contract != nil {
			continue
		}

		contract, err = contractsQ.Insert(data.Contract{
			Name:    list[i].Name,
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

func prepareIndexerParams(ctx context.Context, cfg config.Config) (*newIndexerParams, error) {
	clients, err := helpers.GetNetworkClients(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init network clients")
	}

	certIntegrators, err := helpers.GetCertIntegratorContracts(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init cert integrator contracts")
	}

	issuerBlocks, issuerAddresses, err := updAndGetContractsInfo(
		postgres.NewContractsQ(cfg.DB()),
		cfg.CertificatesIssuer().List,
		data.Issuer,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update and retrieve certificates issuer contracts info")
	}

	fabricBlocks, fabricAddresses, err := updAndGetContractsInfo(
		postgres.NewContractsQ(cfg.DB()),
		cfg.CertificatesFabric().List,
		data.Fabric,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update and retrieve certificates fabric contracts info")
	}

	return &newIndexerParams{
		cfg:             cfg,
		ctx:             ctx,
		clients:         clients,
		issuerCh:        make(chan string),
		certIntegrators: certIntegrators,
		issuerBlocks:    issuerBlocks,
		issuerAddresses: issuerAddresses,
		fabricBlocks:    fabricBlocks,
		fabricAddresses: fabricAddresses,
	}, nil
}
