package user

import (
	"context"

	"github.com/butterv/kubernetes-sample1-app/app/domain/model"
	appstatus "github.com/butterv/kubernetes-sample1-app/app/domain/service/status"
	pb "github.com/butterv/kubernetes-sample1-app/app/gen/go/v1/user"
	"github.com/butterv/kubernetes-sample1-app/app/presenter"
)

func (s *userService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	con := s.r.NewConnection()
	defer con.Close()

	uID := model.UserID(req.GetUserId())
	u, err := con.User().FindByID(ctx, uID)
	if err != nil {
		// TODO(butterv): output error log
		return nil, appstatus.FailedToGetUser.Err()
	}

	return &pb.GetUserResponse{
		User: presenter.UserToPbUser(u),
	}, nil
}
