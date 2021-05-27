//go:generate mockgen -source=$GOFILE -package=mock_persistence -destination=../../infrastructure/mock/$GOFILE
package repository

import (
	"context"

	"github.com/butterv/kubernetes-sample1-app/app/domain/model"
)

type UserRepositoryAccess interface {
	FindByID(ctx context.Context, id model.UserID) (*model.User, error)
	FindByIDs(ctx context.Context, ids []model.UserID) (model.Users, error)
}

type UserRepositoryModify interface {
	Create(ctx context.Context, id model.UserID, email string) error
	DeleteByID(ctx context.Context, id model.UserID) error
}
