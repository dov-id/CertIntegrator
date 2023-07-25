package postgres

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	participantsTableName          = "participants"
	participantsUsersAddressColumn = participantsTableName + ".user_address"
	participantsContractIdColumn   = participantsTableName + ".contract_id"
)

type ParticipantsQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	updateBuilder sq.UpdateBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewParticipantsQ(db *pgdb.DB) data.Participants {
	return &ParticipantsQ{
		db:            db,
		selectBuilder: sq.Select("*").From(usersTableName),
		updateBuilder: sq.Update(usersTableName),
		deleteBuilder: sq.Delete(usersTableName),
	}
}

func (q ParticipantsQ) New() data.Participants {
	return NewParticipantsQ(q.db.Clone())
}

func (q ParticipantsQ) Insert(participant data.Participant) error {
	query := sq.Insert(participantsTableName).SetMap(structs.Map(participant)).
		Suffix("ON CONFLICT DO NOTHING")

	return q.db.Exec(query)
}

func (q ParticipantsQ) Delete() error {
	var deleted []data.Participant

	err := q.db.Select(&deleted, q.deleteBuilder.Suffix("RETURNING *"))
	if err != nil {
		return err
	}

	if len(deleted) == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (q ParticipantsQ) Get() (*data.Participant, error) {
	var result data.Participant
	err := q.db.Get(&result, q.selectBuilder)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q ParticipantsQ) Select() ([]data.Participant, error) {
	var result []data.Participant
	fmt.Println(q.selectBuilder.MustSql())
	return result, q.db.Select(&result, q.selectBuilder)
}

func (q ParticipantsQ) WithUsers() data.Participants {
	q.selectBuilder = sq.Select().
		Columns(usersAddressColumn, usersPubKeyColumn, participantsUsersAddressColumn, participantsContractIdColumn).
		From(participantsTableName).
		LeftJoin(usersTableName + " ON " + usersAddressColumn + " = " + participantsUsersAddressColumn)

	return q
}

func (q ParticipantsQ) FilterByContractId(contractIds ...int64) data.Participants {
	equal := sq.Eq{participantsContractIdColumn: contractIds}

	q.selectBuilder = q.selectBuilder.Where(equal)
	q.updateBuilder = q.updateBuilder.Where(equal)
	q.deleteBuilder = q.deleteBuilder.Where(equal)

	return q
}
