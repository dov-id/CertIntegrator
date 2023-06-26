package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type AttemptsCfg struct {
	Daily int64 `figure:"daily,required"`
}

func (c *config) Attempts() *AttemptsCfg {
	return c.attempts.Do(func() interface{} {
		var cfg AttemptsCfg

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "attempts")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out attempts config"))
		}

		return &cfg
	}).(*AttemptsCfg)
}
