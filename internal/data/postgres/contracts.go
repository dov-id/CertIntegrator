package postgres

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	contractsTableName     = "contracts"
	contractsAddressColumn = contractsTableName + ".address"
)

type ContractsQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewContractsQ(db *pgdb.DB) data.Contracts {
	return &ContractsQ{
		db:            db,
		selectBuilder: sq.Select("*").From(contractsTableName),
		deleteBuilder: sq.Delete(contractsTableName),
	}
}

func (r ContractsQ) New() data.Contracts {
	return NewContractsQ(r.db)
}

func (r ContractsQ) Get() (*data.Contract, error) {
	var result data.Contract
	err := r.db.Get(&result, r.selectBuilder)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (r ContractsQ) Select() ([]data.Contract, error) {
	var result []data.Contract

	err := r.db.Select(&result, r.selectBuilder)

	return result, err
}

func (r ContractsQ) Insert(link data.Contract) (*data.Contract, error) {
	var result data.Contract
	insertStmt := sq.Insert(contractsTableName).SetMap(structs.Map(link)).Suffix("RETURNING *")

	err := r.db.Get(&result, insertStmt)

	return &result, err
}

func (r ContractsQ) Delete() error {
	var deleted []data.Contract

	err := r.db.Select(&deleted, r.deleteBuilder.Suffix("RETURNING *"))
	if err != nil {
		return err
	}

	if len(deleted) == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r ContractsQ) FilterByAddresses(addresses ...string) data.Contracts {
	equalAddresses := sq.Eq{contractsAddressColumn: addresses}

	r.selectBuilder = r.selectBuilder.Where(equalAddresses)
	r.deleteBuilder = r.deleteBuilder.Where(equalAddresses)

	return r
}
