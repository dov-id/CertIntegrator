package postgres

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	TxStatusesTableName     = "tx_statuses"
	TxStatusesTxIdColumn    = TxStatusesTableName + ".tx_id"
	TxStatusesNetworkColumn = TxStatusesTableName + ".network"
	TxCountNetworkColumn    = "count_network"
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
		selectBuilder: sq.Select("*").From(TxStatusesTableName),
		updateBuilder: sq.Update(TxStatusesTableName),
		deleteBuilder: sq.Delete(TxStatusesTableName),
	}
}

func (q TxStatusesQ) New() data.TxStatuses {
	return NewTxStatusesQ(q.db.Clone())
}

func (q TxStatusesQ) Insert(transaction data.TxStatus) error {
	query := sq.Insert(TxStatusesTableName).SetMap(structs.Map(transaction))

	return q.db.Exec(query)
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

	if err == sql.ErrNoRows {
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
	equal := sq.Eq{TxCountNetworkColumn: networksAmount}

	q.selectBuilder = q.selectBuilder.Where(equal)
	q.updateBuilder = q.updateBuilder.Where(equal)
	q.deleteBuilder = q.deleteBuilder.Where(equal)

	return q
}

func (q TxStatusesQ) WithInnerSelect(selector sq.SelectBuilder, alias string) data.TxStatuses {
	q.selectBuilder = q.selectBuilder.FromSelect(selector, alias)

	return q
}

func (q TxStatusesQ) SelectWithCount() {
	//sq.Expr("COUNT(DISTINCT $1)")
	//sq.Alias("COUNT(DISTINCT network)", "count_network")
	//sq.Select(TxStatusesTxIdColumn, " as count_network").From(TxStatusesTableName).GroupBy(txStatusesTxIdColumn)
}
