package user

import "github.com/google/uuid"

type User struct {
	Id       string
	Name     string
	Password string
	Balance  int64
}

func NewUser(name string, password string, balance int64) *User {
	return &User{
		Id:       uuid.NewString(),
		Name:     name,
		Password: password,
		Balance:  balance,
	}
}

func (u *User) Validate() error {
	if u.Balance < 0 {
		return BalanceNegative
	}

	return nil
}
