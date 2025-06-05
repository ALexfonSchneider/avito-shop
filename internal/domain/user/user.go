package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        string
	Username  string
	Password  string
	Balance   int64
	CreatedAt time.Time
}

func NewUser(name string, password string, balance int64, createdAt time.Time) *User {
	return &User{
		ID:        uuid.NewString(),
		Username:  name,
		Password:  password,
		Balance:   balance,
		CreatedAt: createdAt,
	}
}

func (u *User) Validate() error {
	if u.Balance < 0 {
		return BalanceNegative
	}

	return nil
}
