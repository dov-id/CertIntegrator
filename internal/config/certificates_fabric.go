package config

import (
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (c *config) CertificatesFabric() *ContractsCfg {
	return c.certificatesFabric.Do(func() interface{} {
		var cfg ContractsCfg

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks, ContractHooks).
			From(kv.MustGetStringMap(c.getter, "certificates_fabric")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out certificates fabric config"))
		}

		return &cfg
	}).(*ContractsCfg)
}
