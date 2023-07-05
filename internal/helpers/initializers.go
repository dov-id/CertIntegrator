package helpers

import (
	"fmt"

	"github.com/dov-id/cert-integrator-svc/contracts"
	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func InitNetworkClients(networks map[string]config.Network) (map[string]*ethclient.Client, error) {
	clients := make(map[string]*ethclient.Client)

	infura := networks[data.InfuraNetwork]

	for network, params := range networks {
		if network == data.InfuraNetwork || network == data.MetamaskNetwork {
			continue
		}

		var rawUrl string
		switch network {
		case data.QNetwork:
			rawUrl = params.RpcUrl
		default:
			rawUrl = params.RpcUrl + infura.Key
		}

		client, err := ethclient.Dial(rawUrl)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("failed to make dial connect to `%s` network", network))
		}
		clients[network] = client
	}

	return clients, nil
}

func InitCertIntegratorContracts(certIntegrators map[string]string, clients map[string]*ethclient.Client) (map[string]*contracts.CertIntegratorContract, error) {
	certIntegratorContracts := make(map[string]*contracts.CertIntegratorContract)

	for network, address := range certIntegrators {
		contract, err := contracts.NewCertIntegratorContract(common.HexToAddress(address), clients[network])
		if err != nil {
			return nil, errors.Wrap(err, "failed to create new ethereum cert integrator contract")
		}

		certIntegratorContracts[network] = contract
	}

	return certIntegratorContracts, nil
}
