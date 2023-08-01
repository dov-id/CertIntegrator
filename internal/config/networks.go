package config

import (
	"crypto/ecdsa"
	"sync"

	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type NetworksCfg struct {
	Networks map[data.Network]Network
}

type Network struct {
	RpcProviderWsUrl      string
	BlockExplorerApiUrl   string
	BlockExplorerApiKey   string
	CertIntegratorAddress common.Address
	WalletCfg             *WalletCfg
}

type WalletCfg struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

func (c *config) Networks() *NetworksCfg {
	return c.networks.Do(func() interface{} {
		var cfg NetworksCfg

		err := figure.
			Out(&cfg).
			With(NetworkHooks).
			From(kv.MustGetStringMap(c.getter, "networks")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out networks config"))
		}

		return &cfg
	}).(*NetworksCfg)
}

func getKeys(private string) (privateKey *ecdsa.PrivateKey, fromAddress common.Address, err error) {
	var once sync.Once

	once.Do(func() {
		privateKey, err = crypto.HexToECDSA(private)
		if err != nil {
			err = errors.Wrap(err, "failed to convert hex to ecdsa")
			return
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			err = data.ErrFailedToCastKey
			return
		}

		fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
		return
	})

	return privateKey, fromAddress, err
}
