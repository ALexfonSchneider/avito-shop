package transaction

import "errors"

var (
	CannotSendCoinsToYourSelf   = errors.New("Cannot send coins to your self")
	AmountMustBeGreaterThanZero = errors.New("Amount must be greater than zero")
)
