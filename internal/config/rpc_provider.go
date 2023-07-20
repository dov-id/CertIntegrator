package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type RpcProviderCfg struct {
	ApiKey string `figure:"api_key,required"`
}

func (c *config) RpcProvider() *RpcProviderCfg {
	return c.rpcProvider.Do(func() interface{} {
		var cfg RpcProviderCfg

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "rpc_provider")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out rpc provider config"))
		}

		return &cfg
	}).(*RpcProviderCfg)
}
