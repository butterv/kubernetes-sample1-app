package user

import (
	"context"

	"github.com/butterv/kubernetes-sample1-app/app/domain/repository"
	appstatus "github.com/butterv/kubernetes-sample1-app/app/domain/service/status"
	pb "github.com/butterv/kubernetes-sample1-app/app/gen/go/v1/user"
)

func (s *userService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	con := s.r.NewConnection()
	defer con.Close()

	uID := s.userIDGenerator.Generate()
	err := con.RunTransaction(func(tx repository.Transaction) error {
		err := tx.User().Create(ctx, uID, req.GetEmail())
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		// TODO(butterv): output error log
		return nil, appstatus.FailedToCreateUser.Err()
	}

	return &pb.CreateUserResponse{
		UserId: string(uID),
	}, nil
}
