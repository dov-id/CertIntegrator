package data

type ContractAddresses interface {
	New() ContractAddresses

	Insert(address Address) error
	Delete() error
	Get() (*Address, error)
	Select() ([]Address, error)

	FilterByAddresses(addresses ...string) ContractAddresses
}

type Address struct {
	CourseName string `json:"course_name" db:"course_name" structs:"course_name"`
	Address    string `json:"address" db:"address" structs:"address"`
}
