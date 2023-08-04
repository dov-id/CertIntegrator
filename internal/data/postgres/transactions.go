package postgres

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/fatih/structs"
	pkgErrors "github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	transactionTableName     = "transactions"
	transactionsStatusColumn = transactionTableName + ".status"
	transactionsIdColumn     = transactionTableName + ".id"
)

type TransactionsQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	updateBuilder sq.UpdateBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewTransactionsQ(db *pgdb.DB) data.Transactions {
	return &TransactionsQ{
		db:            db,
		selectBuilder: sq.Select("*").From(transactionTableName),
		updateBuilder: sq.Update(transactionTableName),
		deleteBuilder: sq.Delete(transactionTableName),
	}
}

func (q TransactionsQ) New() data.Transactions {
	return NewTransactionsQ(q.db.Clone())
}

func (q TransactionsQ) Insert(transaction data.Transaction) error {
	return q.db.Exec(
		sq.Insert(transactionTableName).
			SetMap(structs.Map(transaction)),
	)
}

func (q TransactionsQ) Delete() error {
	var deleted []data.Transaction

	err := q.db.Select(&deleted, q.deleteBuilder.Suffix("RETURNING *"))
	if err != nil {
		return err
	}

	if len(deleted) == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (q TransactionsQ) Get() (*data.Transaction, error) {
	var result data.Transaction
	err := q.db.Get(&result, q.selectBuilder)

	if pkgErrors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &result, err
}

func (q TransactionsQ) Select() ([]data.Transaction, error) {
	var result []data.Transaction

	return result, q.db.Select(&result, q.selectBuilder)
}

func (q TransactionsQ) Update(transaction data.TransactionToUpdate) error {
	q.updateBuilder = q.updateBuilder.SetMap(structs.Map(transaction))

	return q.db.Exec(q.updateBuilder)
}

func (q TransactionsQ) FilterByStatuses(statuses ...data.TransactionStatus) data.Transactions {
	equal := sq.Eq{transactionsStatusColumn: statuses}

	q.selectBuilder = q.selectBuilder.Where(equal)
	q.updateBuilder = q.updateBuilder.Where(equal)
	q.deleteBuilder = q.deleteBuilder.Where(equal)

	return q
}

func (q TransactionsQ) FilterByIds(ids ...int64) data.Transactions {
	equal := sq.Eq{transactionsIdColumn: ids}

	q.selectBuilder = q.selectBuilder.Where(equal)
	q.updateBuilder = q.updateBuilder.Where(equal)
	q.deleteBuilder = q.deleteBuilder.Where(equal)

	return q
}
