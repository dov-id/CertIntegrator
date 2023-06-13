package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type NetworksCfg struct {
	Networks map[string]Network
}
type Network struct {
	RpcUrl     string
	PrivateKey string
}

type networksCfg struct {
	List []network
}

type network struct {
	Name       string `fig:"name,required"`
	RpcUrl     string `fig:"rpc_url,required"`
	PrivateKey string `fig:"private_key,required"`
}

func (c *config) Networks() *NetworksCfg {
	return c.networks.Do(func() interface{} {
		var cfg networksCfg

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks, NetworkHooks).
			From(kv.MustGetStringMap(c.getter, "networks")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out networks config"))
		}

		mapCfg := createMapNetworks(cfg.List)
		return &mapCfg
	}).(*NetworksCfg)
}

func createMapNetworks(list []network) NetworksCfg {
	var cfg NetworksCfg
	cfg.Networks = make(map[string]Network)

	for _, elem := range list {
		cfg.Networks[elem.Name] = Network{
			RpcUrl:     elem.RpcUrl,
			PrivateKey: elem.PrivateKey,
		}
	}

	return cfg
}
