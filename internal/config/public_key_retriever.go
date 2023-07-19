package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type PublicKeyRetrieverCfg struct {
	DailyAttemptsCount int64 `figure:"daily_attempts_count,required"`
}

func (c *config) PublicKeyRetriever() *PublicKeyRetrieverCfg {
	return c.attempts.Do(func() interface{} {
		var cfg PublicKeyRetrieverCfg

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "public_key_retriever")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out public key retriever config"))
		}

		return &cfg
	}).(*PublicKeyRetrieverCfg)
}
