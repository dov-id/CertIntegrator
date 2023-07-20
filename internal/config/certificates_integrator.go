package config

import (
	"github.com/dov-id/cert-integrator-svc/internal/types"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CertificatesIntegratorCfg struct {
	Addresses map[types.Network]string
}

type integratorsCfg struct {
	List []integrator
}

type integrator struct {
	Network string `fig:"network,required"`
	Address string `fig:"address,required"`
}

func (c *config) CertificatesIntegrator() *CertificatesIntegratorCfg {
	return c.certificatesIntegrator.Do(func() interface{} {
		var cfg integratorsCfg

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks, IntegratorHooks).
			From(kv.MustGetStringMap(c.getter, "certificates_integrator")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out cert integrators config"))
		}

		return createMapIntegrators(cfg.List)
	}).(*CertificatesIntegratorCfg)
}

func createMapIntegrators(list []integrator) *CertificatesIntegratorCfg {
	var cfg CertificatesIntegratorCfg
	cfg.Addresses = make(map[types.Network]string)

	for _, elem := range list {
		cfg.Addresses[types.Network(elem.Network)] = elem.Address
	}

	return &cfg
}
