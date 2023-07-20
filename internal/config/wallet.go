package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type WalletCfg struct {
	PrivateKey string `figure:"private_key,required"`
}

func (c *config) Wallet() *WalletCfg {
	return c.wallet.Do(func() interface{} {
		var cfg WalletCfg

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "wallet")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out wallet config"))
		}

		return &cfg
	}).(*WalletCfg)
}
