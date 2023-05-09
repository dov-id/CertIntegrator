package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CertificatesCfg struct {
	Address       string   `figure:"address"`
	FromBlock     int64    `figure:"from_block"`
	ContractsList []string `figure:"contracts_list"`
}

func (c *config) CertificatesIssuer() *CertificatesCfg {
	return c.certificatesFabric.Do(func() interface{} {
		var cfg CertificatesCfg

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "certificates_issuer")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out certificates issuer config"))
		}

		return &cfg
	}).(*CertificatesCfg)
}
