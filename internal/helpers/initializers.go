package helpers

import (
	"fmt"

	"github.com/dov-id/cert-integrator-svc/contracts"
	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func InitNetworkClients(networks map[data.Network]config.Network) (map[data.Network]*ethclient.Client, error) {
	clients := make(map[data.Network]*ethclient.Client)

	for network, params := range networks {
		client, err := ethclient.Dial(params.RpcProviderWsUrl)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("failed to make dial connect to `%s` network", network))
		}
		clients[network] = client
	}

	return clients, nil
}

func InitCertIntegratorContracts(
	networks map[data.Network]config.Network,
	clients map[data.Network]*ethclient.Client,
) (map[data.Network]*contracts.CertIntegratorContract, error) {
	certIntegratorContracts := make(map[data.Network]*contracts.CertIntegratorContract)

	for network, params := range networks {
		contract, err := contracts.NewCertIntegratorContract(params.CertIntegratorAddress, clients[network])
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("failed to create new `%s` cert integrator contract", network))
		}

		certIntegratorContracts[network] = contract
	}

	return certIntegratorContracts, nil
}
