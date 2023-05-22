package service

import (
	"context"

	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/service/handlers"
	"github.com/dov-id/cert-integrator-svc/internal/service/listeners"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()
	ctx := context.Background()
	s.startListener(ctx, s.cfg)

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
		),
	)
	r.Route("/integrations/cert-integrator-svc", func(r chi.Router) {
		// configure endpoints here
	})

	return r
}

func (s *service) startListener(ctx context.Context, cfg config.Config) {
	conn, err := pgx.Connect(ctx, cfg.DBConfig().URL)
	if err != nil {
		panic(errors.Wrap(err, "failed to connect to db"))
	}

	s.log.Info("Starting fabrics listener")

	for _, contract := range cfg.CertificatesFabric().List {
		listeners.NewListener(
			cfg,
			ctx,
			conn,
			contract.Address,
			contract.FromBlock,
		).Run(ctx)
	}

	s.log.Info("Starting issuers listener")

	for _, contract := range cfg.CertificatesIssuer().List {
		listeners.NewListener(
			cfg,
			ctx,
			conn,
			contract.Address,
			contract.FromBlock,
		).Run(ctx)
	}
}
