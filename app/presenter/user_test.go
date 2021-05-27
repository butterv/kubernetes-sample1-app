package presenter_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	"github.com/butterv/kubernetes-sample1-app/app/domain/model"
	pb "github.com/butterv/kubernetes-sample1-app/app/gen/go/v1/user"
	"github.com/butterv/kubernetes-sample1-app/app/presenter"
)

func TestUserToPbUser(t *testing.T) {
	want := &pb.User{
		UserId: "TEST_ID",
		Email:  "TEST_EMAIL",
	}

	u := &model.User{
		ID:    "TEST_ID",
		Email: "TEST_EMAIL",
	}

	got := presenter.UserToPbUser(u)
	if diff := cmp.Diff(got, want, protocmp.Transform()); diff != "" {
		t.Errorf("presenter.UserToPbUser(%v) = %#v; want %v\ndiff = %s", u, got, want, diff)
	}
}

func TestUserToPbUser_ReturnsNil(t *testing.T) {
	var u *model.User

	got := presenter.UserToPbUser(u)
	if got != nil {
		t.Errorf("presenter.UserToPbUser(%v) = %#v; want nil", u, got)
	}
}

func TestUsersToPbUsers(t *testing.T) {
	want := []*pb.User{
		{
			UserId: "TEST_ID1",
			Email:  "TEST_EMAIL1",
		},
		{
			UserId: "TEST_ID2",
			Email:  "TEST_EMAIL2",
		},
		{
			UserId: "TEST_ID3",
			Email:  "TEST_EMAIL3",
		},
	}

	us := []*model.User{
		{
			ID:    "TEST_ID1",
			Email: "TEST_EMAIL1",
		},
		{
			ID:    "TEST_ID2",
			Email: "TEST_EMAIL2",
		},
		{
			ID:    "TEST_ID3",
			Email: "TEST_EMAIL3",
		},
	}

	got := presenter.UsersToPbUsers(us)
	if diff := cmp.Diff(got, want, protocmp.Transform()); diff != "" {
		t.Errorf("presenter.UsersToPbUsers(%v) = %#v; want %v\ndiff = %s", us, got, want, diff)
	}
}

func TestUsersToPbUsers_ReturnsNil(t *testing.T) {
	var us []*model.User

	got := presenter.UsersToPbUsers(us)
	if got != nil {
		t.Errorf("presenter.UsersToPbUsers(%v) = %#v; want nil", us, got)
	}
}
