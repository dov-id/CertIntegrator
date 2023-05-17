package data

type Contracts interface {
	New() Contracts

	Insert(address Contract) (*Contract, error)
	Delete() error
	Get() (*Contract, error)
	Select() ([]Contract, error)

	FilterByAddresses(addresses ...string) Contracts
}

type Contract struct {
	Id      uint64 `json:"id" db:"id" structs:"-"`
	Name    string `json:"name" db:"name" structs:"name"`
	Address string `json:"address" db:"address" structs:"address"`
}
