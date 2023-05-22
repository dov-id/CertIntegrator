package data

type Contracts interface {
	New() Contracts

	Insert(contract Contract) (*Contract, error)
	Update(contract ContractToUpdate) error
	Delete() error
	Get() (*Contract, error)
	Select() ([]Contract, error)

	FilterByAddresses(addresses ...string) Contracts
}

type Contract struct {
	Id      uint64 `json:"id" db:"id" structs:"-"`
	Name    string `json:"name" db:"name" structs:"name"`
	Address string `json:"address" db:"address" structs:"address"`
	Block   int64  `json:"block" db:"block" structs:"block"`
}

type ContractToUpdate struct {
	Name    *string `structs:"name,omitempty"`
	Address *string `structs:"address,omitempty"`
	Block   *int64  `structs:"block,omitempty"`
}
