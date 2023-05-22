package postgres

import (
	"context"
	"database/sql"

	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/iden3/go-merkletree-sql/v2"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const upsertStmt = `INSERT INTO mt_nodes (mt_id, key, type, child_l, child_r, entry) VALUES ($1, $2, $3, $4, $5, $6) ` +
	`ON CONFLICT (mt_id, key) DO UPDATE SET type = $3, child_l = $4, child_r = $5, entry = $6`

const updateRootStmt = `INSERT INTO mt_roots (mt_id, key) VALUES ($1, $2) ` +
	`ON CONFLICT (mt_id) DO UPDATE SET key = $2`

const getRootStmt = `SELECT mt_id, key, created_at, deleted_at FROM mt_roots WHERE mt_id = $1`

const getKeyStmt = `SELECT mt_id, key, type, child_l, child_r, entry, created_at, deleted_at FROM mt_nodes WHERE mt_id = $1 AND key = $2`

// Storage implements the db.Storage interface
type Storage struct {
	db             *pgdb.DB
	mtId           uint64
	currentVersion uint64
	currentRoot    *merkletree.Hash
}

func NewPGDBStorage(db *pgdb.DB, mtId uint64) *Storage {
	return &Storage{db: db, mtId: mtId}
}

func (s *Storage) Get(ctx context.Context, key []byte) (*merkletree.Node, error) {
	item := data.NodeItem{}
	err := s.db.GetRaw(&item, getKeyStmt, s.mtId, key)

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

func (s *Storage) Put(ctx context.Context, key []byte,
	node *merkletree.Node) error {

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

	err := s.db.ExecRaw(upsertStmt, s.mtId, key[:], node.Type,
		childL, childR, entry)
	return err
}

func (s *Storage) GetRoot(ctx context.Context) (*merkletree.Hash, error) {
	var root merkletree.Hash
	var err error

	if s.currentRoot != nil {
		copy(root[:], s.currentRoot[:])
		return &root, nil
	}

	item := data.RootItem{}
	err = s.db.GetRaw(&item, getRootStmt, s.mtId)
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
	err := s.db.ExecRaw(updateRootStmt, s.mtId, s.currentRoot[:])
	if err != nil {
		err = errors.Wrap(err, "failed to update current root hash")
	}
	return err
}
