package user

import (
	"github.com/butterv/kubernetes-sample1-app/app/domain/model"
	"github.com/butterv/kubernetes-sample1-app/app/domain/repository"
	pb "github.com/butterv/kubernetes-sample1-app/app/gen/go/v1/user"
)

type userService struct {
	r repository.Repository

	userIDGenerator model.UserIDGenerator
}

// NewUserService generates the `UsersServer` implementation.
func NewUserService(r repository.Repository, userIDGenerator model.UserIDGenerator) pb.UsersServer {
	return &userService{
		r:               r,
		userIDGenerator: userIDGenerator,
	}
}
