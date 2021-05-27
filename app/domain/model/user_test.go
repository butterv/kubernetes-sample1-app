package model_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/butterv/kubernetes-sample1-app/app/domain/model"
)

func TestNewDefaultUserIDGenerator(t *testing.T) {
	generator := model.NewDefaultUserIDGenerator()

	_, ok := generator.(model.UserIDGenerator)
	if !ok {
		t.Error("UserIDGenerator is not implemented")
	}
}

func TestDefaultUserIDGenerator_Generate(t *testing.T) {
	generator := model.NewDefaultUserIDGenerator()
	v := generator.Generate()
	i := interface{}(v)

	switch i.(type) {
	case model.UserID:
		// pass
	default:
		t.Errorf("generator.Generate() = %T; want model.UserID", v)
	}
}

func TestUsers_IDs(t *testing.T) {
	want := []model.UserID{"TEST_ID1", "TEST_ID2", "TEST_ID3"}

	now := time.Now()
	us := model.Users{
		{
			ID:        "TEST_ID1",
			Email:     "TEST_EMAIL1",
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
		},
		{
			ID:        "TEST_ID2",
			Email:     "TEST_EMAIL2",
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
		},
		{
			ID:        "TEST_ID3",
			Email:     "TEST_EMAIL3",
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
		},
	}

	got := us.IDs()
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("us.IDs() = %#v; want %v\ndiff = %s", got, want, diff)
	}
}

func TestUsers_FindByID(t *testing.T) {
	now := time.Now()
	want := &model.User{
		ID:        "TEST_ID1",
		Email:     "TEST_EMAIL1",
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	}

	us := model.Users{
		{
			ID:        "TEST_ID1",
			Email:     "TEST_EMAIL1",
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
		},
		{
			ID:        "TEST_ID2",
			Email:     "TEST_EMAIL2",
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
		},
		{
			ID:        "TEST_ID3",
			Email:     "TEST_EMAIL3",
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
		},
	}

	id := model.UserID("TEST_ID1")
	got := us.FindByID(id)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("us.FindByID(%s) = %#v, _; want %v\ndiff = %s", id, got, want, diff)
	}
}

func TestUsers_FindByID_NotFound(t *testing.T) {
	now := time.Now()
	us := model.Users{
		{
			ID:        "TEST_ID1",
			Email:     "TEST_EMAIL1",
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
		},
		{
			ID:        "TEST_ID2",
			Email:     "TEST_EMAIL2",
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
		},
		{
			ID:        "TEST_ID3",
			Email:     "TEST_EMAIL3",
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
		},
	}

	id := model.UserID("TEST_ID4")
	got := us.FindByID(id)
	if got != nil {
		t.Errorf("us.FindByID(%s) = %#v, _; want nil", id, got)
	}
}
