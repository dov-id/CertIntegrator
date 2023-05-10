package postgres

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/dov-id/CertIntegrator/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	contractAddressesTableName     = "contract_addresses"
	contractAddressesAddressColumn = contractAddressesTableName + ".address"
)

type ContractAddressesQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewContractAddressesQ(db *pgdb.DB) data.ContractAddresses {
	return &ContractAddressesQ{
		db:            db,
		selectBuilder: sq.Select("*").From(contractAddressesTableName),
		deleteBuilder: sq.Delete(contractAddressesTableName),
	}
}

func (r ContractAddressesQ) New() data.ContractAddresses {
	return NewContractAddressesQ(r.db)
}

func (r ContractAddressesQ) Get() (*data.Address, error) {
	var result data.Address
	err := r.db.Get(&result, r.selectBuilder)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (r ContractAddressesQ) Select() ([]data.Address, error) {
	var result []data.Address

	err := r.db.Select(&result, r.selectBuilder)

	return result, err
}

func (r ContractAddressesQ) Insert(link data.Address) error {
	insertStmt := sq.Insert(contractAddressesTableName).SetMap(structs.Map(link)).Suffix("ON CONFLICT (address) DO NOTHING")

	return r.db.Exec(insertStmt)
}

func (r ContractAddressesQ) Delete() error {
	var deleted []data.Address

	err := r.db.Select(&deleted, r.deleteBuilder.Suffix("RETURNING *"))
	if err != nil {
		return err
	}

	if len(deleted) == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r ContractAddressesQ) FilterByAddresses(addresses ...string) data.ContractAddresses {
	equalAddresses := sq.Eq{contractAddressesAddressColumn: addresses}

	r.selectBuilder = r.selectBuilder.Where(equalAddresses)
	r.deleteBuilder = r.deleteBuilder.Where(equalAddresses)

	return r
}
