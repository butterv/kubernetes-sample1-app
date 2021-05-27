package user

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/butterv/kubernetes-sample1-app/app/domain/model"
)

type userRepositoryModify struct {
	tx *sqlx.Tx
}

func NewUserRepositoryModify(tx *sqlx.Tx) *userRepositoryModify {
	if tx == nil {
		return nil
	}

	return &userRepositoryModify{
		tx: tx,
	}
}

func (r *userRepositoryModify) Create(ctx context.Context, id model.UserID, email string) error {
	u := &model.User{
		ID:    id,
		Email: email,
	}

	_, err := r.tx.NamedExecContext(ctx, "INSERT INTO users (id, email) VALUES (:id, :email)", u)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryModify) DeleteByID(ctx context.Context, id model.UserID) error {
	_, err := r.tx.ExecContext(ctx, "UPDATE users SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
