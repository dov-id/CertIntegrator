package service

import (
	"context"

	"github.com/dov-id/CertIntegrator/internal/config"
	"github.com/dov-id/CertIntegrator/internal/service/handlers"
	"github.com/dov-id/CertIntegrator/internal/service/listeners"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
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
	r.Route("/integrations/CertIntegrator", func(r chi.Router) {
		// configure endpoints here
	})

	return r
}

func (s *service) startListener(ctx context.Context, cfg config.Config) {
	s.log.Info("Starting issuer listener")

	listeners.NewListener(
		cfg.Infura().Key,
		cfg.CertificatesIssuer().Address,
		cfg.CertificatesIssuer().FromBlock,
	).Run(ctx, cfg)

	s.log.Info("Starting fabric listener")

	listeners.NewListener(
		cfg.Infura().Key,
		cfg.CertificatesFabric().Address,
		cfg.CertificatesFabric().FromBlock,
	).Run(ctx, cfg)
}
