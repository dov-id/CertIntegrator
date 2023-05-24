package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CertificatesIntegratorCfg struct {
	Ethereum string `figure:"ethereum,required"`
	Polygon  string `figure:"polygon,required"`
	Q        string `figure:"q,required"`
}

func (c *config) CertificatesIntegrator() *CertificatesIntegratorCfg {
	return c.certificatesIntegrator.Do(func() interface{} {
		var cfg CertificatesIntegratorCfg

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "certificates_integrator")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out certificates integrator config"))
		}

		return &cfg
	}).(*CertificatesIntegratorCfg)
}
