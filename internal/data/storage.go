package data

import "github.com/iden3/go-merkletree-sql/v2"

type NodeItem struct {
	MTId      uint64  `db:"mt_id"`
	Key       []byte  `db:"key"`
	Type      byte    `db:"type"`
	ChildL    []byte  `db:"child_l"`
	ChildR    []byte  `db:"child_r"`
	Entry     []byte  `db:"entry"`
	CreatedAt *uint64 `db:"created_at"`
	DeletedAt *uint64 `db:"deleted_at"`
}

type RootItem struct {
	MTId      uint64  `db:"mt_id"`
	Key       []byte  `db:"key"`
	CreatedAt *uint64 `db:"created_at"`
	DeletedAt *uint64 `db:"deleted_at"`
}

func (item *NodeItem) Node() (*merkletree.Node, error) {
	node := merkletree.Node{
		Type: merkletree.NodeType(item.Type),
	}
	if item.ChildL != nil {
		node.ChildL = &merkletree.Hash{}
		copy(node.ChildL[:], item.ChildL[:])
	}
	if item.ChildR != nil {
		node.ChildR = &merkletree.Hash{}
		copy(node.ChildR[:], item.ChildR[:])
	}
	if len(item.Entry) > 0 {
		if len(item.Entry) != 2*merkletree.ElemBytesLen {
			return nil, merkletree.ErrNodeBytesBadSize
		}
		node.Entry = [2]*merkletree.Hash{{}, {}}
		copy(node.Entry[0][:], item.Entry[0:32])
		copy(node.Entry[1][:], item.Entry[32:64])
	}
	return &node, nil
}
