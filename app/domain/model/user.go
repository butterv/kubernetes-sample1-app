package model

import (
	"time"

	"github.com/rs/xid"
)

type UserID string

// UserIDGenerator is an user ID generator interface.
type UserIDGenerator interface {
	Generate() UserID
}

type defaultUserIDGenerator struct{}

// NewDefaultUserIDGenerator generates a default user id generator.
func NewDefaultUserIDGenerator() UserIDGenerator {
	return &defaultUserIDGenerator{}
}

func (*defaultUserIDGenerator) Generate() UserID {
	return UserID(xid.New().String())
}

type User struct {
	ID        UserID     `db:"id"`
	Email     string     `db:"email"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type Users []*User

func (us Users) IDs() []UserID {
	uIDs := make([]UserID, 0, len(us))

	for _, u := range us {
		uIDs = append(uIDs, u.ID)
	}

	return uIDs
}

func (us Users) FindByID(id UserID) *User {
	for _, u := range us {
		if u.ID == id {
			return u
		}
	}

	return nil
}
