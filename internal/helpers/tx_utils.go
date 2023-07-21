package helpers

import (
	"context"
	"math/big"

	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func GetAuth(ctx context.Context, client *ethclient.Client, walletCfg *config.WalletCfg) (*bind.TransactOpts, error) {
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain id")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(walletCfg.PrivateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transaction signer")
	}

	nonce, err := client.PendingNonceAt(ctx, walletCfg.Address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get nonce")
	}

	auth.Nonce = big.NewInt(int64(nonce))

	return auth, nil
}

func WaitForTransactionMined(client *ethclient.Client, transaction *types.Transaction, ctx context.Context, log *logan.Entry) {
	go func() {
		log.WithField("tx", transaction.Hash().Hex()).Debugf("waiting to mine")

		_, err := bind.WaitMined(ctx, client, transaction)
		if err != nil {
			panic(errors.Wrap(err, "failed to mine transaction"))
		}

		log.WithField("tx", transaction.Hash().Hex()).Debugf("was mined")
	}()
}
