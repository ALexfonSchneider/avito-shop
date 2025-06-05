package transaction

import (
	"errors"
	"github.com/ALexfonSchneider/avito-shop/internal/domain"
)

var (
	CannotSendCoinsToYourSelf   = domain.Error{Err: errors.New("cannot send coins to your self")}
	AmountMustBeGreaterThanZero = domain.Error{Err: errors.New("amount must be greater than zero")}
)
