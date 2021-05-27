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

func TestUserService_GetUser(t *testing.T) {
	uID := model.UserID("TEST_USER_ID")
	email := "TEST_EMAIL"

	want := &pb.GetUserResponse{
		User: &pb.User{
			UserId: string(uID),
			Email:  email,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock_persistence.New(ctrl)
	r.UserRepositoryAccess.EXPECT().
		FindByID(gomock.Any(), uID).
		DoAndReturn(func(context.Context, model.UserID) (*model.User, error) {
			now := time.Now()
			return &model.User{
				ID:        uID,
				Email:     email,
				CreatedAt: now,
				UpdatedAt: now,
			}, nil
		})

	service := user.NewUserService(r, model.NewDefaultUserIDGenerator())

	ctx := context.Background()
	req := &pb.GetUserRequest{
		UserId: string(uID),
	}

	got, err := service.GetUser(ctx, req)
	if err != nil {
		t.Fatalf("service.GetUser(ctx, %v) = _, %#v; want nil", req, err)
	}
	if diff := cmp.Diff(got, want, protocmp.Transform()); diff != "" {
		t.Errorf("service.GetUser(ctx, %v) = %#v, _; want %v\ndiff = %s", req, got, want, diff)
	}
}

func TestUserService_GetUser_NotFound(t *testing.T) {
	want := &pb.GetUserResponse{}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uID := model.UserID("TEST_USER_ID")

	r := mock_persistence.New(ctrl)
	r.UserRepositoryAccess.EXPECT().
		FindByID(gomock.Any(), uID).
		DoAndReturn(func(context.Context, model.UserID) (*model.User, error) {
			return nil, nil
		})

	service := user.NewUserService(r, model.NewDefaultUserIDGenerator())

	ctx := context.Background()
	req := &pb.GetUserRequest{
		UserId: string(uID),
	}

	got, err := service.GetUser(ctx, req)
	if err != nil {
		t.Fatalf("service.GetUser(ctx, %v) = _, %#v; want nil", req, err)
	}
	if diff := cmp.Diff(got, want, protocmp.Transform()); diff != "" {
		t.Errorf("service.GetUser(ctx, %v) = %#v, _; want %v\ndiff = %s", req, got, want, diff)
	}
}

func TestUserService_GetUser_FindByIDReturnsError(t *testing.T) {
	wantErr := appstatus.FailedToGetUser.Err()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uID := model.UserID("TEST_USER_ID")

	r := mock_persistence.New(ctrl)
	r.UserRepositoryAccess.EXPECT().
		FindByID(gomock.Any(), uID).
		DoAndReturn(func(context.Context, model.UserID) (*model.User, error) {
			return nil, errors.New("an error occurred")
		})

	service := user.NewUserService(r, model.NewDefaultUserIDGenerator())

	ctx := context.Background()
	req := &pb.GetUserRequest{
		UserId: string(uID),
	}

	_, err := service.GetUser(ctx, req)
	if err == nil {
		t.Fatalf("service.GetUser(ctx, %v) = _, nil; want %v", req, wantErr)
	}
	if !errors.Is(err, wantErr) {
		t.Errorf("service.GetUser(ctx, %v) = _, %v; want %v", req, err, wantErr)
	}
}
