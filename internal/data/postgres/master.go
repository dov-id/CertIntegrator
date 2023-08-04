package postgres

import (
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type MasterQ struct {
	db *pgdb.DB
}

func NewMasterQ(db *pgdb.DB) data.MasterQ {
	return &MasterQ{
		db: db.Clone(),
	}
}

func (q *MasterQ) New() data.MasterQ {
	return NewMasterQ(q.db)
}

func (q *MasterQ) UsersQ() data.Users {
	return NewUsersQ(q.db)
}

func (q *MasterQ) ContractsQ() data.Contracts {
	return NewContractsQ(q.db)
}

func (q *MasterQ) TransactionsQ() data.Transactions {
	return NewTransactionsQ(q.db)
}

func (q *MasterQ) TxStatusesQ() data.TxStatuses {
	return NewTxStatusesQ(q.db)
}

func (q *MasterQ) Transaction(fn func() error) error {
	return q.db.Transaction(func() error {
		return fn()
	})
}
