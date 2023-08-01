package data

type MasterQ interface {
	New() MasterQ

	UsersQ() Users
	ContractsQ() Contracts

	TransactionsQ() Transactions
	TxStatusesQ() TxStatuses

	Transaction(func() error) error
}
