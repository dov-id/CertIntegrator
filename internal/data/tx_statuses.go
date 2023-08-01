package data

type TxStatuses interface {
	New() TxStatuses

	Insert(txStatus TxStatus) error
	Delete() error
	Get() (*TxStatus, error)
	Select() ([]TxStatus, error)

	WithCountNetworkColumn() TxStatuses
	FilterByNetworksAmount(networksAmount ...int64) TxStatuses
}

type TxStatus struct {
	TxId         int64  `json:"tx_id" db:"tx_id" structs:"tx_id"`
	Network      string `json:"network" db:"network" structs:"network"`
	CountNetwork string `json:",omitempty" db:"count_network" structs:",omitempty"`
}
