package data

import sq "github.com/Masterminds/squirrel"

type TxStatuses interface {
	New() TxStatuses

	Insert(txStatus TxStatus) error
	Delete() error
	Get() (*TxStatus, error)
	Select() ([]TxStatus, error)

	WithInnerSelect(selector sq.SelectBuilder, alias string) TxStatuses
	FilterByNetworksAmount(networksAmount ...int64) TxStatuses
}

type TxStatus struct {
	TxId         int64  `json:"tx_id" db:"tx_id" structs:"tx_id"`
	Network      string `json:"network" db:"network" structs:"network"`
	CountNetwork string `json:",omitempty" db:"count_network" structs:",omitempty"`
}
