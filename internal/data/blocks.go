package data

type Blocks interface {
	New() Blocks

	Upsert(block Block) error
	Delete() error
	Get() (*Block, error)

	FilterByContractNames(contractNames ...string) Blocks
}

type Block struct {
	ContractName    string `json:"contract_name" db:"contract_name" structs:"contract_name"`
	LastBlockNumber int64  `json:"last_block_number" db:"last_block_number" structs:"last_block_number"`
}
