package postgres

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/fatih/structs"
	pkgErrors "github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	txStatusesTableName     = "tx_statuses"
	txStatusesTxIdColumn    = txStatusesTableName + ".tx_id"
	txStatusesNetworkColumn = txStatusesTableName + ".network"
	txCountNetworkColumn    = "count_network"
)

type TxStatusesQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	updateBuilder sq.UpdateBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewTxStatusesQ(db *pgdb.DB) data.TxStatuses {
	return &TxStatusesQ{
		db:            db,
		selectBuilder: sq.Select("*").From(txStatusesTableName),
		updateBuilder: sq.Update(txStatusesTableName),
		deleteBuilder: sq.Delete(txStatusesTableName),
	}
}

func (q TxStatusesQ) New() data.TxStatuses {
	return NewTxStatusesQ(q.db.Clone())
}

func (q TxStatusesQ) Insert(transaction data.TxStatus) error {
	return q.db.Exec(
		sq.Insert(txStatusesTableName).
			SetMap(structs.Map(transaction)),
	)
}

func (q TxStatusesQ) Delete() error {
	var deleted []data.TxStatus

	err := q.db.Select(&deleted, q.deleteBuilder.Suffix("RETURNING *"))
	if err != nil {
		return err
	}

	if len(deleted) == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (q TxStatusesQ) Get() (*data.TxStatus, error) {
	var result data.TxStatus
	err := q.db.Get(&result, q.selectBuilder)

	if pkgErrors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &result, err
}

func (q TxStatusesQ) Select() ([]data.TxStatus, error) {
	var result []data.TxStatus

	return result, q.db.Select(&result, q.selectBuilder)
}

func (q TxStatusesQ) GroupBy(columns ...string) data.TxStatuses {
	q.selectBuilder = q.selectBuilder.GroupBy(columns...)

	return q
}

func (q TxStatusesQ) FilterByNetworksAmount(networksAmount ...int64) data.TxStatuses {
	equal := sq.Eq{txCountNetworkColumn: networksAmount}

	q.selectBuilder = q.selectBuilder.Where(equal)
	q.updateBuilder = q.updateBuilder.Where(equal)
	q.deleteBuilder = q.deleteBuilder.Where(equal)

	return q
}

func (q TxStatusesQ) WithCountNetworkColumn() data.TxStatuses {
	innerSelect := sq.Select().
		Columns(txStatusesTxIdColumn).
		Column(sq.Alias(sq.Expr(fmt.Sprintf("COUNT(DISTINCT %s)", txStatusesNetworkColumn)), txCountNetworkColumn)).
		From(txStatusesTableName).
		GroupBy(txStatusesTxIdColumn)

	q.selectBuilder = q.selectBuilder.FromSelect(innerSelect, "t")

	return q
}
