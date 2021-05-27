package user

import (
	"context"

	"github.com/butterv/kubernetes-sample1-app/app/domain/model"
	appstatus "github.com/butterv/kubernetes-sample1-app/app/domain/service/status"
	pb "github.com/butterv/kubernetes-sample1-app/app/gen/go/v1/user"
	"github.com/butterv/kubernetes-sample1-app/app/presenter"
)

func (s *userService) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	con := s.r.NewConnection()
	defer con.Close()

	var uIDs []model.UserID
	for _, uID := range req.GetUserIds() {
		uIDs = append(uIDs, model.UserID(uID))
	}

	us, err := con.User().FindByIDs(ctx, uIDs)
	if err != nil {
		// TODO(butterv): output error log
		return nil, appstatus.FailedToGetUsers.Err()
	}

	return &pb.GetUsersResponse{
		Users: presenter.UsersToPbUsers(us),
	}, nil
}
