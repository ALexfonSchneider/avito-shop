package user

import "errors"

var (
	BalanceNegative = errors.New("balance cannot be negative")
)
