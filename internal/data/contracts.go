package data

type ContractType string

const (
	FABRIC ContractType = "fabric"
	ISSUER ContractType = "issuer"
)

type Contracts interface {
	New() Contracts

	Insert(contract Contract) (*Contract, error)
	Update(contract ContractToUpdate) error
	Delete() error
	Get() (*Contract, error)
	Select() ([]Contract, error)

	FilterByAddresses(addresses ...string) Contracts
	FilterByTypes(types ...ContractType) Contracts
}

type Contract struct {
	Id      uint64       `json:"id" db:"id" structs:"-"`
	Name    string       `json:"name" db:"name" structs:"name"`
	Address string       `json:"address" db:"address" structs:"address"`
	Block   int64        `json:"block" db:"block" structs:"block"`
	Type    ContractType `json:"type" db:"type" structs:"type"`
}

type ContractToUpdate struct {
	Name    *string `structs:"name,omitempty"`
	Address *string `structs:"address,omitempty"`
	Block   *int64  `structs:"block,omitempty"`
	Type    *string `structs:"type,omitempty"`
}
