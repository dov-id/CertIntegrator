package indexer

import (
	"context"

	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/helpers"
	"gitlab.com/distributed_lab/logan/v3"
)

type IIndexer interface {
	Run(ctx context.Context)
}

type Indexer struct {
	cfg config.Config
	ctx context.Context
	log *logan.Entry

	Addresses  []string
	Blocks     []int64
	ContractsQ data.Contracts

	Cancel context.CancelFunc
}

func Run(cfg config.Config, ctx context.Context) {
	cancelCtx, cancelFn := context.WithCancel(ctx)
	blocks, addresses := helpers.SeparateContractArrays(cfg.CertificatesIssuer().List)
	NewIndexer(
		cfg,
		cancelCtx,
		addresses,
		blocks,
		nil,
	).Run(ctx)

	blocks, addresses = helpers.SeparateContractArrays(cfg.CertificatesFabric().List)
	NewIndexer(
		cfg,
		ctx,
		addresses,
		blocks,
		cancelFn,
	).Run(ctx)
}
