package postgres

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/iden3/go-merkletree-sql/v2"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	mtRootsTable = "mt_roots"
	mtNodesTable = "mt_nodes"

	mtIdColumn      = "mt_id"
	keyColumn       = "key"
	typeColumn      = "type"
	childLColumn    = "child_l"
	childRColumn    = "child_r"
	entryColumn     = "entry"
	createdAtColumn = "created_at"
	deletedAtColumn = "deleted_at"
)

// DeleteMTree deletes merkle tree info from storage by current id
func (s *Storage) DeleteMTree(ctx context.Context) error {
	query := sq.Delete(mtRootsTable).Where(sq.Eq{mtIdColumn: s.mtId})
	err := s.db.Exec(query)
	if err != nil {
		return errors.Wrap(err, "failed to delete merkle tree root")
	}

	query = sq.Delete(mtNodesTable).Where(sq.Eq{mtIdColumn: s.mtId})
	err = s.db.Exec(query)
	if err != nil {
		return errors.Wrap(err, "failed to delete merkle tree nodes")
	}

	return nil
}

// Storage implements the db.Storage interface
type Storage struct {
	db             *pgdb.DB
	mtId           uint64
	currentVersion uint64
	currentRoot    *merkletree.Hash
}

func NewStorage(db *pgdb.DB, mtId uint64) *Storage {
	return &Storage{db: db, mtId: mtId}
}

func (s *Storage) Get(ctx context.Context, key []byte) (*merkletree.Node, error) {
	item := data.NodeItem{}

	/*
		SELECT mt_id, key, type, child_l, child_r, entry, created_at, deleted_at FROM mt_nodes WHERE mt_id = $1 AND key = $2
	*/

	stmt := sq.
		Select(mtIdColumn, keyColumn, typeColumn, childLColumn, childRColumn, entryColumn, createdAtColumn, deletedAtColumn).
		From(mtNodesTable).
		Where(sq.Eq{mtIdColumn: s.mtId, keyColumn: key})

	err := s.db.Get(&item, stmt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, merkletree.ErrNotFound
		}
		return nil, err
	}

	node, err := item.Node()
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (s *Storage) Put(ctx context.Context, key []byte, node *merkletree.Node) error {
	var childL []byte
	if node.ChildL != nil {
		childL = append(childL, node.ChildL[:]...)
	}

	var childR []byte
	if node.ChildR != nil {
		childR = append(childR, node.ChildR[:]...)
	}

	var entry []byte
	if node.Entry[0] != nil && node.Entry[1] != nil {
		entry = append(node.Entry[0][:], node.Entry[1][:]...)
	}

	/*
		`INSERT INTO mt_nodes (mt_id, key, type, child_l, child_r, entry) VALUES ($1, $2, $3, $4, $5, $6) ` +
			`ON CONFLICT (mt_id, key) DO UPDATE SET type = $3, child_l = $4, child_r = $5, entry = $6`
	*/

	updateStmt, args := sq.Update(" ").
		Set("type", node.Type).
		Set("child_l", childL).
		Set("child_r", childR).
		Set("entry", entry).
		MustSql()

	query := sq.Insert(mtNodesTable).
		Columns(mtIdColumn, keyColumn, typeColumn, childLColumn, childRColumn, entryColumn).
		Values(s.mtId, key[:], node.Type, childL, childR, entry).
		Suffix("ON CONFLICT (mt_id, key) DO "+updateStmt, args...)

	return s.db.Exec(query)
}

func (s *Storage) GetRoot(ctx context.Context) (*merkletree.Hash, error) {
	var root merkletree.Hash
	var err error

	if s.currentRoot != nil {
		copy(root[:], s.currentRoot[:])
		return &root, nil
	}

	/*
		`SELECT mt_id, key, created_at, deleted_at FROM mt_roots WHERE mt_id = $1`
	*/

	item := data.RootItem{}

	stmt := sq.Select(mtIdColumn, keyColumn, createdAtColumn, deletedAtColumn).
		From(mtRootsTable).
		Where(sq.Eq{mtIdColumn: s.mtId})

	err = s.db.Get(&item, stmt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, merkletree.ErrNotFound
		}
		return nil, err
	}

	if s.currentRoot == nil {
		s.currentRoot = &merkletree.Hash{}
	}
	copy(s.currentRoot[:], item.Key[:])
	copy(root[:], s.currentRoot[:])
	return &root, nil
}

func (s *Storage) SetRoot(ctx context.Context, hash *merkletree.Hash) error {
	if s.currentRoot == nil {
		s.currentRoot = &merkletree.Hash{}
	}
	copy(s.currentRoot[:], hash[:])

	/*
		`INSERT INTO mt_roots (mt_id, key) VALUES ($1, $2) ` +
			`ON CONFLICT (mt_id) DO UPDATE SET key = $2`
	*/

	updateQuery, args := sq.Update(" ").
		Set("key", s.currentRoot[:]).
		MustSql()

	query := sq.
		Insert(mtRootsTable).
		Columns(mtIdColumn, keyColumn).
		Values(s.mtId, s.currentRoot[:]).
		Suffix("ON CONFLICT (mt_id) DO "+updateQuery, args...)

	err := s.db.Exec(query)
	if err != nil {
		err = errors.Wrap(err, "failed to update current root hash")
	}
	return err
}
