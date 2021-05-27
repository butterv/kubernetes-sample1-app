package persistence

import (
	"github.com/jmoiron/sqlx"

	"github.com/butterv/kubernetes-sample1-app/app/domain/repository"
	"github.com/butterv/kubernetes-sample1-app/app/infrastructure/persistence/user"
)

type persistenceRepository struct {
	db *sqlx.DB
}

type persistenceConnection struct {
	db *sqlx.DB
}

type persistenceTransaction struct {
	tx *sqlx.Tx
}

func New(db *sqlx.DB) repository.Repository {
	return &persistenceRepository{
		db: db,
	}
}

func (r *persistenceRepository) NewConnection() repository.Connection {
	return &persistenceConnection{
		db: r.db,
	}
}

func (c *persistenceConnection) Close() error {
	// We don't need to close *sqlx.DB. No need to do anything.
	return nil
}

func (c *persistenceConnection) RunTransaction(f func(repository.Transaction) error) error {
	tx, err := c.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()

			panic(p) // Re-throw panic
		}
	}()

	err = f(&persistenceTransaction{
		tx: tx,
	})
	if err != nil {
		_ = tx.Rollback()

		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c *persistenceConnection) User() repository.UserRepositoryAccess {
	return user.NewUserRepositoryAccess(c.db)
}

func (c *persistenceTransaction) User() repository.UserRepositoryModify {
	return user.NewUserRepositoryModify(c.tx)
}
