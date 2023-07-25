package sender

import (
	"context"

	"github.com/dov-id/cert-integrator-svc/contracts"
	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"github.com/dov-id/cert-integrator-svc/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
)

type Sender interface {
	Run(ctx context.Context)
}

type sender struct {
	cfg config.Config
	log *logan.Entry

	TransactionsQ data.Transactions
	TxStatusesQ   data.TxStatuses

	Clients         map[types.Network]*ethclient.Client
	CertIntegrators map[types.Network]*contracts.CertIntegratorContract
}

type updateStateParams struct {
	network        types.Network
	client         *ethclient.Client
	ids            []int64
	courses        []common.Address
	states         [][32]byte
	certIntegrator *contracts.CertIntegratorContract
}

//func Run(cfg config.Config, ctx context.Context) {
//	NewSender(cfg, _, _).Run(ctx)
//}

func NewSender(cfg config.Config, clients map[types.Network]*ethclient.Client, certIntegrators map[types.Network]*contracts.CertIntegratorContract) Sender {
	return &sender{
		cfg:             cfg,
		log:             cfg.Log(),
		TransactionsQ:   postgres.NewTransactionsQ(cfg.DB().Clone()),
		TxStatusesQ:     postgres.NewTxStatusesQ(cfg.DB().Clone()),
		Clients:         clients,
		CertIntegrators: certIntegrators,
	}
}
