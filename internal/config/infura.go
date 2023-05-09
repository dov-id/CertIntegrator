package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type InfuraCfg struct {
	Key string `fig:"key,required"`
}

func (c *config) Infura() *InfuraCfg {
	return c.infura.Do(func() interface{} {
		var cfg InfuraCfg

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "infura")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out infura config"))
		}

		return &cfg
	}).(*InfuraCfg)
}
