package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	"github.com/butterv/kubernetes-sample1-app/app/domain/model"
	appstatus "github.com/butterv/kubernetes-sample1-app/app/domain/service/status"
	"github.com/butterv/kubernetes-sample1-app/app/domain/service/user"
	pb "github.com/butterv/kubernetes-sample1-app/app/gen/go/v1/user"
	mock_persistence "github.com/butterv/kubernetes-sample1-app/app/infrastructure/mock"
)

type testUserIDGenerator struct {
	userID model.UserID
}

func newTestUserIDGenerator(uID model.UserID) model.UserIDGenerator {
	return &testUserIDGenerator{
		userID: uID,
	}
}

func (g *testUserIDGenerator) Generate() model.UserID {
	return g.userID
}

func TestUserService_CreateUser(t *testing.T) {
	uID := model.UserID("TEST_USER_ID")
	email := "TEST_EMAIL"

	want := &pb.CreateUserResponse{
		UserId: string(uID),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock_persistence.New(ctrl)
	r.UserRepositoryModify.EXPECT().
		Create(gomock.Any(), uID, email).
		DoAndReturn(func(context.Context, model.UserID, string) error {
			return nil
		})

	userIDGenerator := newTestUserIDGenerator(uID)
	service := user.NewUserService(r, userIDGenerator)

	ctx := context.Background()
	req := &pb.CreateUserRequest{
		Email: email,
	}

	got, err := service.CreateUser(ctx, req)
	if err != nil {
		t.Fatalf("service.CreateUser(ctx, %v) = _, %#v; want nil", req, err)
	}
	if diff := cmp.Diff(got, want, protocmp.Transform()); diff != "" {
		t.Errorf("service.CreateUser(ctx, %v) = %#v, _; want %v\ndiff = %s", req, got, want, diff)
	}
}

func TestUserService_CreateUser_CreateReturnsError(t *testing.T) {
	wantErr := appstatus.FailedToCreateUser.Err()

	uID := model.UserID("TEST_USER_ID")
	email := "TEST_EMAIL"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock_persistence.New(ctrl)
	r.UserRepositoryModify.EXPECT().
		Create(gomock.Any(), uID, email).
		DoAndReturn(func(context.Context, model.UserID, string) error {
			return errors.New("an error occurred")
		})

	userIDGenerator := newTestUserIDGenerator(uID)
	service := user.NewUserService(r, userIDGenerator)

	ctx := context.Background()
	req := &pb.CreateUserRequest{
		Email: email,
	}

	_, err := service.CreateUser(ctx, req)
	if err == nil {
		t.Fatalf("service.CreateUser(ctx, %v) = _, nil; want %v", req, wantErr)
	}
	if !errors.Is(err, wantErr) {
		t.Errorf("service.CreateUser(ctx, %v) = _, %v; want %v", req, err, wantErr)
	}
}
