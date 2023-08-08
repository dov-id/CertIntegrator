package data

type Users interface {
	New() Users

	Insert(user User) error
	Delete() error
	Get() (*User, error)
	Select() ([]User, error)

	FilterByAddresses(addresses ...string) Users
	FilterByContractId(contractIds ...uint64) Users

	Limit(num uint64) Users
	Offset(num uint64) Users
}

type User struct {
	Address    string `json:"address" db:"address" structs:"address"`
	ContractId uint64 `json:"contract_id" db:"contract_id" structs:"contract_id"`
}
