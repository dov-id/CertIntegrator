package config

import (
	"github.com/dov-id/cert-integrator-svc/internal/types"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type NetworksCfg struct {
	Networks map[types.Network]Network
}
type Network struct {
	RpcUrl      string
	HttpsApiUrl string
	ApiKey      string
}

type networksCfg struct {
	List []network
}

type network struct {
	Name        string `fig:"name,required"`
	RpcUrl      string `fig:"rpc_url,required"`
	HttpsApiUrl string `fig:"https_api_url,required"`
	ApiKey      string `fig:"api_key,required"`
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

		return createMapNetworks(cfg.List)
	}).(*NetworksCfg)
}

func createMapNetworks(list []network) *NetworksCfg {
	var cfg NetworksCfg
	cfg.Networks = make(map[types.Network]Network)

	for _, elem := range list {
		cfg.Networks[types.Network(elem.Name)] = Network{
			RpcUrl:      elem.RpcUrl,
			HttpsApiUrl: elem.HttpsApiUrl,
			ApiKey:      elem.ApiKey,
		}
	}

	return &cfg
}
