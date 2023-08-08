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
	usersTableName        = "users"
	usersAddressColumn    = usersTableName + ".address"
	usersPubKeyColumn     = usersTableName + ".public_key"
	usersContractIdColumn = usersTableName + ".contract_id"
)

type UsersQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	updateBuilder sq.UpdateBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewUsersQ(db *pgdb.DB) data.Users {
	return &UsersQ{
		db:            db,
		selectBuilder: sq.Select("*").From(usersTableName),
		updateBuilder: sq.Update(usersTableName),
		deleteBuilder: sq.Delete(usersTableName),
	}
}

func (q UsersQ) New() data.Users {
	return NewUsersQ(q.db.Clone())
}

func (q UsersQ) Get() (*data.User, error) {
	var result data.User
	err := q.db.Get(&result, q.selectBuilder)

	if pkgErrors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &result, err
}

func (q UsersQ) Select() ([]data.User, error) {
	var result []data.User

	return result, q.db.Select(&result, q.selectBuilder)
}

func (q UsersQ) Insert(user data.User) error {
	return q.db.Exec(
		sq.Insert(usersTableName).
			SetMap(structs.Map(user)).
			Suffix("ON CONFLICT (address, contract_id) DO NOTHING"),
	)
}

func (q UsersQ) Delete() error {
	var deleted []data.User

	err := q.db.Select(&deleted, q.deleteBuilder.Suffix("RETURNING *"))
	if err != nil {
		return err
	}

	if len(deleted) == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (q UsersQ) FilterByAddresses(addresses ...string) data.Users {
	equalAddresses := sq.Eq{usersAddressColumn: addresses}

	q.selectBuilder = q.selectBuilder.Where(equalAddresses)
	q.updateBuilder = q.updateBuilder.Where(equalAddresses)
	q.deleteBuilder = q.deleteBuilder.Where(equalAddresses)

	return q
}

func (q UsersQ) FilterByContractId(contractIds ...uint64) data.Users {
	equal := sq.Eq{usersContractIdColumn: contractIds}

	q.selectBuilder = q.selectBuilder.Where(equal)
	q.updateBuilder = q.updateBuilder.Where(equal)
	q.deleteBuilder = q.deleteBuilder.Where(equal)

	return q
}

func (q UsersQ) Limit(num uint64) data.Users {
	q.selectBuilder = q.selectBuilder.Limit(num)
	q.updateBuilder = q.updateBuilder.Limit(num)
	q.deleteBuilder = q.deleteBuilder.Limit(num)

	return q
}

func (q UsersQ) Offset(num uint64) data.Users {
	q.selectBuilder = q.selectBuilder.Offset(num)
	q.updateBuilder = q.updateBuilder.Offset(num)
	q.deleteBuilder = q.deleteBuilder.Offset(num)

	return q
}
