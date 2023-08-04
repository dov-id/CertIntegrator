package cleaner

import (
	"context"

	"github.com/dov-id/cert-integrator-svc/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

const (
	serviceName = "cleaner"
)

func (c *cleaner) Run(ctx context.Context) {
	go running.WithBackOff(
		ctx,
		c.log,
		serviceName,
		c.listen,
		c.cfg.Timeouts().Cleaner,
		c.cfg.Timeouts().Cleaner,
		c.cfg.Timeouts().Cleaner,
	)
}

func (c *cleaner) listen(ctx context.Context) error {
	c.log.Debugf("start cleaner")

	err := c.processHandledTransactions()
	if err != nil {
		return errors.Wrap(err, "failed to process handled transactions")
	}

	c.log.Debugf("stop cleaner")
	return nil
}

func (c *cleaner) processHandledTransactions() error {
	txStatuses, err := c.MasterQ.TxStatusesQ().
		WithCountNetworkColumn().
		FilterByNetworksAmount(data.NetworksAmount).
		Select()
	if err != nil {
		return errors.Wrap(err, "failed to select transactions statuses")
	}

	for _, txStatus := range txStatuses {
		err = c.MasterQ.TransactionsQ().FilterByIds(txStatus.TxId).Delete()
		if err != nil {
			return errors.Wrap(err, "failed to delete mined transactions")
		}
	}

	return nil
}
