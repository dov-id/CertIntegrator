package data

type TransactionStatus string

const (
	PENDING     TransactionStatus = "pending"
	IN_PROGRESS TransactionStatus = "in progress"
)

type Transactions interface {
	New() Transactions

	Insert(transaction Transaction) error
	Update(transaction TransactionToUpdate) error
	Delete() error
	Get() (*Transaction, error)
	Select() ([]Transaction, error)

	FilterByStatuses(status ...TransactionStatus) Transactions
	FilterByIds(ids ...int64) Transactions
}

type Transaction struct {
	Id     int64             `json:"id" db:"id" structs:"-"`
	Status TransactionStatus `json:"status" db:"status" structs:"status"`
	Course string            `json:"course" db:"course" structs:"course"`
	State  []byte            `json:"state" db:"state" structs:"state"`
}

type TransactionToUpdate struct {
	Status *TransactionStatus `structs:"status,omitempty"`
}
