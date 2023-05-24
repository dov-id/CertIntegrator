package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type MetamaskCfg struct {
	PrivateKey string `figure:"private_key,required"`
}

func (c *config) Metamask() *MetamaskCfg {
	return c.metamask.Do(func() interface{} {
		var cfg MetamaskCfg

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "metamask")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out metamask config"))
		}

		return &cfg
	}).(*MetamaskCfg)
}
