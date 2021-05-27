package user_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	"github.com/butterv/kubernetes-sample1-app/app/domain/model"
	appstatus "github.com/butterv/kubernetes-sample1-app/app/domain/service/status"
	"github.com/butterv/kubernetes-sample1-app/app/domain/service/user"
	pb "github.com/butterv/kubernetes-sample1-app/app/gen/go/v1/user"
	mock_persistence "github.com/butterv/kubernetes-sample1-app/app/infrastructure/mock"
)

func TestUserService_GetUsers(t *testing.T) {
	uIDs := []model.UserID{"TEST_USER_ID1", "TEST_USER_ID2", "TEST_USER_ID3"}
	emails := []string{"TEST_EMAIL1", "TEST_EMAIL2", "TEST_EMAIL3"}

	want := &pb.GetUsersResponse{
		Users: []*pb.User{
			{
				UserId: string(uIDs[0]),
				Email:  emails[0],
			},
			{
				UserId: string(uIDs[1]),
				Email:  emails[1],
			},
			{
				UserId: string(uIDs[2]),
				Email:  emails[2],
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock_persistence.New(ctrl)
	r.UserRepositoryAccess.EXPECT().
		FindByIDs(gomock.Any(), uIDs).
		DoAndReturn(func(context.Context, []model.UserID) (model.Users, error) {
			now := time.Now()
			return model.Users{
				{
					ID:        uIDs[0],
					Email:     emails[0],
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					ID:        uIDs[1],
					Email:     emails[1],
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					ID:        uIDs[2],
					Email:     emails[2],
					CreatedAt: now,
					UpdatedAt: now,
				},
			}, nil
		})

	service := user.NewUserService(r, model.NewDefaultUserIDGenerator())

	ctx := context.Background()
	req := &pb.GetUsersRequest{
		UserIds: []string{string(uIDs[0]), string(uIDs[1]), string(uIDs[2])},
	}

	got, err := service.GetUsers(ctx, req)
	if err != nil {
		t.Fatalf("service.GetUsers(ctx, %v) = _, %#v; want nil", req, err)
	}
	if diff := cmp.Diff(got, want, protocmp.Transform()); diff != "" {
		t.Errorf("service.GetUsers(ctx, %v) = %#v, _; want %v\ndiff = %s", req, got, want, diff)
	}
}

func TestUserService_GetUsers_NotFound(t *testing.T) {
	want := &pb.GetUsersResponse{}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uIDs := []model.UserID{"TEST_USER_ID1", "TEST_USER_ID2", "TEST_USER_ID3"}

	r := mock_persistence.New(ctrl)
	r.UserRepositoryAccess.EXPECT().
		FindByIDs(gomock.Any(), uIDs).
		DoAndReturn(func(context.Context, []model.UserID) (model.Users, error) {
			return nil, nil
		})

	service := user.NewUserService(r, model.NewDefaultUserIDGenerator())

	ctx := context.Background()
	req := &pb.GetUsersRequest{
		UserIds: []string{string(uIDs[0]), string(uIDs[1]), string(uIDs[2])},
	}

	got, err := service.GetUsers(ctx, req)
	if err != nil {
		t.Fatalf("service.GetUsers(ctx, %v) = _, %#v; want nil", req, err)
	}
	if diff := cmp.Diff(got, want, protocmp.Transform()); diff != "" {
		t.Errorf("service.GetUsers(ctx, %v) = %#v, _; want %v\ndiff = %s", req, got, want, diff)
	}
}

func TestUserService_GetUsers_FindByIDsReturnsError(t *testing.T) {
	wantErr := appstatus.FailedToGetUsers.Err()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uIDs := []model.UserID{"TEST_USER_ID1", "TEST_USER_ID2", "TEST_USER_ID3"}

	r := mock_persistence.New(ctrl)
	r.UserRepositoryAccess.EXPECT().
		FindByIDs(gomock.Any(), uIDs).
		DoAndReturn(func(context.Context, []model.UserID) (model.Users, error) {
			return nil, errors.New("an error occurred")
		})

	service := user.NewUserService(r, model.NewDefaultUserIDGenerator())

	ctx := context.Background()
	req := &pb.GetUsersRequest{
		UserIds: []string{string(uIDs[0]), string(uIDs[1]), string(uIDs[2])},
	}

	_, err := service.GetUsers(ctx, req)
	if err == nil {
		t.Fatalf("service.GetUsers(ctx, %v) = _, nil; want %v", req, wantErr)
	}
	if !errors.Is(err, wantErr) {
		t.Errorf("service.GetUsers(ctx, %v) = _, %v; want %v", req, err, wantErr)
	}
}
