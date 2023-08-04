package service

import (
	"context"
	"sync"

	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/helpers"
	"github.com/dov-id/cert-integrator-svc/internal/service/api"
	"github.com/dov-id/cert-integrator-svc/internal/service/cleaner"
	"github.com/dov-id/cert-integrator-svc/internal/service/indexer"
	"github.com/dov-id/cert-integrator-svc/internal/service/sender"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Runner = func(config config.Config, context context.Context)

var availableServices = map[string]Runner{
	"api":     api.Run,
	"indexer": indexer.Run,
	"cleaner": cleaner.Run,
	"sender":  sender.Run,
}

func Run(cfg config.Config) {
	logger := cfg.Log().WithField("service", "main")
	wg := new(sync.WaitGroup)
	//sender.NewSender(cfg, params.clients, params.certIntegrators).Run(ctx)

	ctx, err := prepareContextStorage(context.Background(), cfg)
	if err != nil {
		panic(errors.Wrap(err, "failed to prepare context storage"))
	}

	logger.Debugf("Starting all available services")

	for serviceName, service := range availableServices {
		wg.Add(1)

		go func(name string, runner Runner) {
			defer wg.Done()

			runner(cfg, ctx)

		}(serviceName, service)

		logger.WithField("service", serviceName).Debugf("Service started")
	}

	wg.Wait()
}

func prepareContextStorage(ctx context.Context, cfg config.Config) (context.Context, error) {
	clients, err := helpers.InitNetworkClients(cfg.Networks().Networks)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize network clients")
	}

	ctx = context.WithValue(ctx, data.NetworkClients, clients)

	certIntegrators, err := helpers.InitCertIntegratorContracts(cfg.Networks().Networks, clients)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize cert integrator contracts")
	}

	ctx = context.WithValue(ctx, data.CertIntegratorContracts, certIntegrators)

	return ctx, nil
}
