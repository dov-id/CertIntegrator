package sender

import (
	"context"

	"github.com/dov-id/cert-integrator-svc/contracts"
	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"github.com/dov-id/cert-integrator-svc/internal/helpers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"
)

type Sender interface {
	Run(ctx context.Context)
}

type sender struct {
	cfg config.Config
	log *logan.Entry

	MasterQ data.MasterQ

	Clients         map[data.Network]*ethclient.Client
	CertIntegrators map[data.Network]*contracts.CertIntegratorContract
}

type updateStateParams struct {
	network        data.Network
	client         *ethclient.Client
	ids            []int64
	courses        []common.Address
	states         [][32]byte
	certIntegrator *contracts.CertIntegratorContract
}

func Run(cfg config.Config, ctx context.Context) {
	NewSender(ctx, cfg).Run(ctx)
}

func NewSender(ctx context.Context, cfg config.Config) Sender {
	clients, err := helpers.GetNetworkClients(ctx)
	if err != nil {
		panic(errors.Wrap(err, "failed to get network clients"))
	}

	certIntegrators, err := helpers.GetCertIntegratorContracts(ctx)
	if err != nil {
		panic(errors.Wrap(err, "failed to get cert integrator contracts"))
	}

	return &sender{
		cfg:             cfg,
		log:             cfg.Log(),
		MasterQ:         postgres.NewMasterQ(cfg.DB().Clone()),
		Clients:         clients,
		CertIntegrators: certIntegrators,
	}
}
