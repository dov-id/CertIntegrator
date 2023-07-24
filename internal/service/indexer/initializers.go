package indexer

import (
	"context"
	"fmt"

	"github.com/dov-id/cert-integrator-svc/contracts"
	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"github.com/dov-id/cert-integrator-svc/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func InitNetworkClients(networks map[types.Network]config.Network, rpcProvider *config.RpcProviderCfg) (map[types.Network]*ethclient.Client, error) {
	clients := make(map[types.Network]*ethclient.Client)

	for network, params := range networks {
		rawUrl := fmt.Sprint(params.RpcUrl, rpcProvider.ApiKey)
		if network == data.QNetwork {
			rawUrl = params.RpcUrl
		}

		client, err := ethclient.Dial(rawUrl)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("failed to make dial connect to `%s` network", network))
		}
		clients[network] = client
	}

	return clients, nil
}

func initCertIntegratorContracts(
	certIntegrators map[types.Network]string,
	clients map[types.Network]*ethclient.Client,
) (map[types.Network]*contracts.CertIntegratorContract, error) {
	certIntegratorContracts := make(map[types.Network]*contracts.CertIntegratorContract)

	for network, address := range certIntegrators {
		contract, err := contracts.NewCertIntegratorContract(common.HexToAddress(address), clients[network])
		if err != nil {
			return nil, errors.Wrap(err, "failed to create new ethereum cert integrator contract")
		}

		certIntegratorContracts[network] = contract
	}

	return certIntegratorContracts, nil
}

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

func prepareIndexerParams(cfg config.Config, ctx context.Context) (*newIndexerParams, error) {
	clients, err := InitNetworkClients(cfg.Networks().Networks, cfg.RpcProvider())
	if err != nil {
		return nil, errors.Wrap(err, "failed to init network clients")
	}

	certIntegrators, err := initCertIntegratorContracts(cfg.CertificatesIntegrator().Addresses, clients)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init cert integrator contracts")
	}

	issuerBlocks, issuerAddresses, err := updAndGetContractsInfo(
		postgres.NewContractsQ(cfg.DB()),
		cfg.CertificatesIssuer().List,
		data.ISSUER,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update and retrieve certificates issuer contracts info")
	}

	fabricBlocks, fabricAddresses, err := updAndGetContractsInfo(
		postgres.NewContractsQ(cfg.DB()),
		cfg.CertificatesFabric().List,
		data.FABRIC,
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
