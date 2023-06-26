package postgres

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	usersTableName     = "users"
	usersAddressColumn = usersTableName + ".address"
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
	return NewUsersQ(q.db)
}

func (q UsersQ) Get() (*data.User, error) {
	var result data.User
	err := q.db.Get(&result, q.selectBuilder)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q UsersQ) Select() ([]data.User, error) {
	var result []data.User

	err := q.db.Select(&result, q.selectBuilder)

	return result, err
}

func (q UsersQ) Upsert(user data.User) error {
	updateStmt, args := sq.Update(" ").
		Set("public_key", user.PublicKey).
		MustSql()

	query := sq.Insert(usersTableName).SetMap(structs.Map(user)).
		Suffix("ON CONFLICT (address) DO "+updateStmt, args...)

	return q.db.Exec(query)
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
