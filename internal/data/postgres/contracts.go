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
	contractsTypeColumn    = contractsTableName + ".type"
)

type ContractsQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	updateBuilder sq.UpdateBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewContractsQ(db *pgdb.DB) data.Contracts {
	return &ContractsQ{
		db:            db,
		selectBuilder: sq.Select("*").From(contractsTableName),
		updateBuilder: sq.Update(contractsTableName),
		deleteBuilder: sq.Delete(contractsTableName),
	}
}

func (r ContractsQ) New() data.Contracts {
	return NewContractsQ(r.db.Clone())
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

	return result, r.db.Select(&result, r.selectBuilder)
}

func (r ContractsQ) Insert(contract data.Contract) (*data.Contract, error) {
	var result data.Contract
	insertStmt := sq.Insert(contractsTableName).SetMap(structs.Map(contract)).Suffix("RETURNING *")

	err := r.db.Get(&result, insertStmt)

	return &result, err
}

func (r ContractsQ) Update(contract data.ContractToUpdate) error {
	r.updateBuilder = r.updateBuilder.SetMap(structs.Map(contract))

	return r.db.Exec(r.updateBuilder)
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
	r.updateBuilder = r.updateBuilder.Where(equalAddresses)
	r.deleteBuilder = r.deleteBuilder.Where(equalAddresses)

	return r
}

func (r ContractsQ) FilterByTypes(types ...data.ContractType) data.Contracts {
	equalTypes := sq.Eq{contractsTypeColumn: types}

	r.selectBuilder = r.selectBuilder.Where(equalTypes)
	r.updateBuilder = r.updateBuilder.Where(equalTypes)
	r.deleteBuilder = r.deleteBuilder.Where(equalTypes)

	return r
}
