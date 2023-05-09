package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (c *config) CertificatesFabric() *CertificatesCfg {
	return c.certificatesFabric.Do(func() interface{} {
		var cfg CertificatesCfg

		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "certificates_fabric")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out certificates fabric config"))
		}

		return &cfg
	}).(*CertificatesCfg)
}
