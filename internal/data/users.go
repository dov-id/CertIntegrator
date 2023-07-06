package data

type Users interface {
	New() Users

	Upsert(user User) error
	Delete() error
	Get() (*User, error)
	Select() ([]User, error)

	FilterByAddresses(addresses ...string) Users
}

type User struct {
	Address   string `json:"address" db:"address" structs:"address"`
	PublicKey string `json:"public_key" db:"public_key" structs:"public_key"`
}
