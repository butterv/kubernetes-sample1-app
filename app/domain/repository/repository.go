package repository

type Repository interface {
	NewConnection() Connection
}

type Connection interface {
	Close() error
	RunTransaction(f func(Transaction) error) error

	User() UserRepositoryAccess
}

type Transaction interface {
	User() UserRepositoryModify
}
