package cleaner

import (
	"context"

	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"gitlab.com/distributed_lab/logan/v3"
)

type Cleaner interface {
	Run(ctx context.Context)
}

type cleaner struct {
	cfg config.Config
	log *logan.Entry

	MasterQ data.MasterQ
}

func Run(cfg config.Config, ctx context.Context) {
	NewCleaner(cfg).Run(ctx)
}

func NewCleaner(cfg config.Config) Cleaner {
	return &cleaner{
		cfg:     cfg,
		log:     cfg.Log(),
		MasterQ: postgres.NewMasterQ(cfg.DB().Clone()),
	}
}
