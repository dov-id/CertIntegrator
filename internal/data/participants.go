package data

type Participants interface {
	New() Participants

	Upsert(participant Participant) error
	Delete() error
	Get() (*Participant, error)
	Select() ([]Participant, error)

	WithUsers() Participants
	FilterByContractId(contractIds ...int64) Participants
}

type Participant struct {
	UserAddress string `json:"user_address" db:"user_address" structs:"user_address"`
	ContractId  int64  `json:"contract_id" db:"contract_id" structs:"contract_id"`
	*User       `structs:",omitempty"`
}
