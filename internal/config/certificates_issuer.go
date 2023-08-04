package config

import (
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Contract struct {
	Address   string `fig:"address,required"`
	Name      string `fig:"name,required"`
	FromBlock int64  `fig:"from_block,required"`
}

type ContractsCfg struct {
	List []Contract
}

func (c *config) CertificatesIssuer() *ContractsCfg {
	return c.certificatesIssuer.Do(func() interface{} {
		var cfg ContractsCfg

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks, ContractHooks).
			From(kv.MustGetStringMap(c.getter, "certificates_issuer")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out certificates issuer config"))
		}

		return &cfg
	}).(*ContractsCfg)
}
