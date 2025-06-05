package user

import (
	"errors"
	"github.com/ALexfonSchneider/avito-shop/internal/domain"
)

var (
	BalanceNegative = domain.Error{Err: errors.New("balance cannot be negative")}
	NotFound        = domain.Error{Err: errors.New("user not found")}
)
