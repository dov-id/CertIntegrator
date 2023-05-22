package service

import (
	"context"
	"sync"

	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/service/api"
	"github.com/dov-id/cert-integrator-svc/internal/service/indexer"
)

type Runner = func(config config.Config, context context.Context)

var availableServices = map[string]Runner{
	"api":     api.Run,
	"indexer": indexer.Run,
}

func Run(cfg config.Config) {
	logger := cfg.Log().WithField("service", "main")
	ctx := context.Background()
	wg := new(sync.WaitGroup)

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
