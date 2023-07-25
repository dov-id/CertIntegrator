package sender

import (
	"context"
	"fmt"
	"math/big"

	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	pkgErrors "github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

const (
	serviceName = "sender"
)

func (s *sender) Run(ctx context.Context) {
	go running.WithBackOff(
		ctx,
		s.log,
		serviceName,
		s.listen,
		s.cfg.Timeouts().Sender,
		s.cfg.Timeouts().Sender,
		s.cfg.Timeouts().Sender,
	)
}

func (s *sender) listen(ctx context.Context) error {
	s.log.Debugf("start sender")

	err := s.processTxs(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to process transactions")
	}

	s.log.Debugf("stop sender")
	return nil
}

func (s *sender) processTxs(ctx context.Context) error {
	txs, err := s.TransactionsQ.FilterByStatuses(data.PENDING).Select()
	if err != nil {
		return errors.Wrap(err, "failed to select transactions")
	}

	txsLength := len(txs)
	if txsLength == 0 {
		s.log.Debugf("no transactions to send were found")
		return nil
	}

	ids := make([]int64, txsLength)
	courses := make([]common.Address, txsLength)
	states := make([][32]byte, txsLength)

	for i := 0; i < txsLength; i++ {
		ids[i] = txs[i].Id
		courses[i] = common.HexToAddress(txs[i].Course)
		copy(states[i][:], txs[i].State[:])
	}

	err = s.publishStates(ctx, updateStateParams{
		states:  states,
		courses: courses,
		ids:     ids,
	})
	if err != nil {
		return errors.Wrap(err, "failed to publish states")
	}

	return nil
}

func (s *sender) publishStates(ctx context.Context, params updateStateParams) error {
	for network, client := range s.Clients {
		params.network = network
		params.client = client
		params.certIntegrator = s.CertIntegrators[network]

		err := s.sendUpdates(ctx, params)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to publishStates in `%s`", network))
		}
	}

	return nil
}

func (s *sender) sendUpdates(ctx context.Context, params updateStateParams) error {
	auth, err := helpers.GetAuth(ctx, params.client, s.cfg.Wallet())
	if err != nil {
		return errors.Wrap(err, "failed to get auth options")
	}

	err = s.sendUpdateCourseState(ctx, auth, params)
	if err != nil {
		return errors.Wrap(err, "failed to update course state")
	}

	return nil
}

func (s *sender) sendUpdateCourseState(ctx context.Context, auth *bind.TransactOpts, params updateStateParams) error {
	transaction, err := params.certIntegrator.UpdateCourseState(auth, params.courses, params.states)
	if err != nil {
		if pkgErrors.Is(err, data.ErrReplacementTxUnderpriced) {
			auth.Nonce = big.NewInt(auth.Nonce.Int64() + 1)
			return s.sendUpdateCourseState(ctx, auth, params)
		}

		return errors.Wrap(err, "failed to update course state")
	}

	s.waitForTransactionMined(ctx, transaction, params)

	return nil
}

func (s *sender) waitForTransactionMined(ctx context.Context, transaction *types.Transaction, params updateStateParams) {
	go func() {
		s.log.WithField("tx", transaction.Hash().Hex()).Debugf("waiting to mine")

		status := data.IN_PROGRESS
		err := s.TransactionsQ.FilterByIds(params.ids...).Update(data.TransactionToUpdate{Status: &status})
		if err != nil {
			panic(errors.Wrap(err, "failed to update tx status"))
		}

		_, err = bind.WaitMined(ctx, params.client, transaction)
		if err != nil {
			panic(errors.Wrap(err, "failed to mine transaction"))
		}

		for _, id := range params.ids {
			err = s.TxStatusesQ.Insert(data.TxStatus{
				TxId:    id,
				Network: params.network.String(),
			})
			if err != nil {
				panic(errors.Wrap(err, "failed to insert tx status"))
			}
		}

		s.log.WithField("tx", transaction.Hash().Hex()).Debugf("was mined")
	}()

}
