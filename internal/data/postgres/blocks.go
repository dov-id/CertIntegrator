package postgres

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/dov-id/CertIntegrator/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	blocksTableName          = "blocks"
	blocksContractNameColumn = blocksTableName + ".contract_name"
)

type BlockQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewBlocksQ(db *pgdb.DB) data.Blocks {
	return &BlockQ{
		db:            db.Clone(),
		selectBuilder: sq.Select("*").From(blocksTableName),
		deleteBuilder: sq.Delete(blocksTableName),
	}
}

func (q BlockQ) New() data.Blocks {
	return NewBlocksQ(q.db)
}

func (q BlockQ) Upsert(block data.Block) error {
	clauses := structs.Map(block)

	updateStmt, args, err := sq.Update(" ").
		Set("block", block.LastBlockNumber).
		ToSql()
	if err != nil {
		return err
	}

	query := sq.Insert(blocksTableName).SetMap(clauses).Suffix("ON CONFLICT (contract_name) DO "+updateStmt, args...)

	return q.db.Exec(query)
}

func (q BlockQ) Delete() error {
	var deleted []data.Block

	err := q.db.Select(&deleted, q.deleteBuilder.Suffix("RETURNING *"))
	if err != nil {
		return err
	}

	if len(deleted) == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (q BlockQ) Get() (*data.Block, error) {
	var result data.Block

	err := q.db.Get(&result, q.selectBuilder)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q BlockQ) FilterByContractNames(names ...string) data.Blocks {
	equalNames := sq.Eq{blocksContractNameColumn: names}

	q.selectBuilder = q.selectBuilder.Where(equalNames)
	q.deleteBuilder = q.deleteBuilder.Where(equalNames)

	return q
}
