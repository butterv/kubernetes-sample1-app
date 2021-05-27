package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/butterv/kubernetes-sample1-app/app/domain/model"
)

type userRepositoryAccess struct {
	db *sqlx.DB
}

func NewUserRepositoryAccess(db *sqlx.DB) *userRepositoryAccess {
	if db == nil {
		return nil
	}

	return &userRepositoryAccess{
		db: db,
	}
}

func (r *userRepositoryAccess) FindByID(ctx context.Context, id model.UserID) (*model.User, error) {
	var u model.User

	err := r.db.GetContext(ctx, &u, "SELECT * FROM users WHERE id = ? AND deleted_at IS NULL", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &u, nil
}

func (r *userRepositoryAccess) FindByIDs(ctx context.Context, ids []model.UserID) (model.Users, error) {
	var us model.Users

	query, params, err := sqlx.In("SELECT * FROM users WHERE id IN (?) AND deleted_at IS NULL", ids)
	if err != nil {
		return nil, err
	}

	err = r.db.SelectContext(ctx, &us, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return us, nil
}
