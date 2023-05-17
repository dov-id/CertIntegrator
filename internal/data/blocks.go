package data

type Blocks interface {
	New() Blocks

	Upsert(block Block) error
	Delete() error
	Get() (*Block, error)

	FilterByContractAddress(contractAddresses ...string) Blocks
}

type Block struct {
	ContractAddress string `json:"contract_address" db:"contract_address" structs:"contract_address"`
	LastBlockNumber int64  `json:"last_block_number" db:"last_block_number" structs:"last_block_number"`
}
