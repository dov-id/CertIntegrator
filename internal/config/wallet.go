package config

import (
	"crypto/ecdsa"
	"sync"

	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type WalletCfg struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

type wallet struct {
	PrivateKey string `figure:"private_key,required"`
}

func (c *config) Wallet() *WalletCfg {
	return c.wallet.Do(func() interface{} {
		var cfg wallet

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "wallet")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out wallet config"))
		}

		return cfg.convert()
	}).(*WalletCfg)
}

func (w *wallet) convert() *WalletCfg {
	privateKey, fromAddress, err := getKeys(w.PrivateKey)
	if err != nil {
		panic(errors.Wrap(err, "failed to get keys"))
	}

	return &WalletCfg{
		PrivateKey: privateKey,
		Address:    fromAddress,
	}
}

func getKeys(private string) (*ecdsa.PrivateKey, common.Address, error) {
	var once sync.Once
	var privateKey *ecdsa.PrivateKey
	var fromAddress common.Address
	var err error

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

	return privateKey, fromAddress, nil
}
