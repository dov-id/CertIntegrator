package data

type TxStatus string

const (
	CREATED     TxStatus = "created"
	IN_PROGRESS TxStatus = "in progress"
)

type Transactions interface {
	New() Transactions

	Insert(transaction Transaction) error
	Update(transaction TransactionToUpdate) error
	Delete() error
	Get() (*Transaction, error)
	Select() ([]Transaction, error)

	FilterByStatuses(status ...TxStatus) Transactions
	FilterByIds(ids ...int64) Transactions
}

type Transaction struct {
	Id     int64    `json:"id" db:"id" structs:"-"`
	Status TxStatus `json:"status" db:"status" structs:"status"`
	Course string   `json:"course" db:"course" structs:"course"`
	State  []byte   `json:"state" db:"state" structs:"state"`
}

type TransactionToUpdate struct {
	Status *TxStatus `structs:"status,omitempty"`
}
