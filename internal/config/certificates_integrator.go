package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CertificatesIntegratorCfg struct {
	Addresses map[string]string
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
			panic(errors.Wrap(err, "failed to figure out networks config"))
		}

		mapCfg := createMapIntegrators(cfg.List)
		return &mapCfg
	}).(*CertificatesIntegratorCfg)
}

func createMapIntegrators(list []integrator) CertificatesIntegratorCfg {
	var cfg CertificatesIntegratorCfg
	cfg.Addresses = make(map[string]string)

	for _, elem := range list {
		cfg.Addresses[elem.Network] = elem.Address
	}

	return cfg
}
