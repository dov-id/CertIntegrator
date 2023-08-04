package helpers

import (
	"context"

	"github.com/dov-id/cert-integrator-svc/contracts"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetNetworkClients(ctx context.Context) (map[data.Network]*ethclient.Client, error) {
	value := ctx.Value(data.NetworkClients)

	clients, ok := value.(map[data.Network]*ethclient.Client)
	if !ok {
		return nil, data.ErrFailedToCastClients
	}

	return clients, nil
}

func GetCertIntegratorContracts(ctx context.Context) (map[data.Network]*contracts.CertIntegratorContract, error) {
	value := ctx.Value(data.CertIntegratorContracts)

	clients, ok := value.(map[data.Network]*contracts.CertIntegratorContract)
	if !ok {
		return nil, data.ErrFailedToCastCertIntegrators
	}

	return clients, nil
}
