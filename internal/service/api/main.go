package api

import (
	"context"
	"net"
	"net/http"

	"github.com/dov-id/cert-integrator-svc/internal/config"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Router struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener

	cfg config.Config
	ctx context.Context
}

func (s *Router) run() error {
	s.log.Debug("Service started")
	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config, ctx context.Context) *Router {
	return &Router{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		cfg:      cfg,
		ctx:      ctx,
	}
}

func Run(cfg config.Config, ctx context.Context) {
	if err := newService(cfg, ctx).run(); err != nil {
		panic(err)
	}
}
