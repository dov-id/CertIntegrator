package postgres

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/dov-id/CertIntegrator-svc/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	blocksTableName             = "blocks"
	blocksContractAddressColumn = blocksTableName + ".contract_address"
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

	updateStmt, args := sq.Update(" ").
		Set("last_block_number", block.LastBlockNumber).
		MustSql()

	query := sq.Insert(blocksTableName).SetMap(clauses).Suffix("ON CONFLICT (contract_address) DO "+updateStmt, args...)

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

func (q BlockQ) FilterByContractAddress(contractAddresses ...string) data.Blocks {
	equalAddresses := sq.Eq{blocksContractAddressColumn: contractAddresses}

	q.selectBuilder = q.selectBuilder.Where(equalAddresses)
	q.deleteBuilder = q.deleteBuilder.Where(equalAddresses)

	return q
}
