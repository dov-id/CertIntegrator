package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type DBCfg struct {
	URL string `fig:"url,required"`
}

func (c *config) DBConfig() *DBCfg {
	return c.db.Do(func() interface{} {
		var cfg DBCfg

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "db")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out database config"))
		}

		return &cfg
	}).(*DBCfg)
}
