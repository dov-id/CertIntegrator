package helpers

import (
	"context"
	"math/big"

	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
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
