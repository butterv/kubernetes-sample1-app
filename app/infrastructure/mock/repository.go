package mock_persistence

import (
	"github.com/golang/mock/gomock"

	"github.com/butterv/kubernetes-sample1-app/app/domain/repository"
)

type mocks struct {
	UserRepositoryAccess *MockUserRepositoryAccess
	UserRepositoryModify *MockUserRepositoryModify
}

type MockRepository struct {
	*mocks
}

type MockConnection struct {
	*mocks
}

type MockTransaction struct {
	*mocks
}

func New(ctrl *gomock.Controller) *MockRepository {
	return &MockRepository{
		mocks: &mocks{
			UserRepositoryAccess: NewMockUserRepositoryAccess(ctrl),
			UserRepositoryModify: NewMockUserRepositoryModify(ctrl),
		},
	}
}

func (r *MockRepository) NewConnection() repository.Connection {
	return &MockConnection{
		mocks: r.mocks,
	}
}

func (c *MockConnection) Close() error {
	return nil
}

func (c *MockConnection) RunTransaction(f func(repository.Transaction) error) error {
	err := f(&MockTransaction{
		mocks: c.mocks,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *MockConnection) User() repository.UserRepositoryAccess {
	return c.UserRepositoryAccess
}

func (t *MockTransaction) User() repository.UserRepositoryModify {
	return t.UserRepositoryModify
}
