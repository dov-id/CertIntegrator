package cleaner

import (
	"context"
	"fmt"

	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
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
	/*
			SELECT * FROM (
				SELECT
					tx_statuses.tx_id,
					(COUNT(DISTINCT tx_statuses.network)) AS count_network
				FROM tx_statuses
				GROUP BY tx_statuses.tx_id
			) AS t WHERE count_network IN (3);

		it gives ability to select only transactions that was mined in 3 network
	*/
	innerSelect := postgres.BuildSelect(postgres.TxStatusesTableName, postgres.TxStatusesTxIdColumn)
	countDistinctColumn := postgres.BuildExpression(fmt.Sprintf("COUNT(DISTINCT %s)", postgres.TxStatusesNetworkColumn))
	innerSelect = postgres.ColumnWithAlias(innerSelect, countDistinctColumn, postgres.TxCountNetworkColumn)
	innerSelect = postgres.GroupBy(innerSelect, postgres.TxStatusesTxIdColumn)

	txStatuses, err := c.TxStatusesQ.WithInnerSelect(innerSelect, "t").FilterByNetworksAmount(data.NetworksAmount).Select()
	if err != nil {
		return errors.Wrap(err, "failed to select transactions statuses")
	}

	for _, txStatus := range txStatuses {
		err = c.TransactionsQ.FilterByIds(txStatus.TxId).Delete()
		if err != nil {
			return errors.Wrap(err, "failed to delete mined transactions")
		}
	}

	return nil
}
