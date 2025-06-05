package application

import (
	"errors"
	"fmt"
)

type Error struct {
	Err error
}

func (e Error) Error() string {
	return fmt.Sprintf("application error: %v", e.Err)
}

var (
	MerchNotFound             = Error{Err: errors.New("merch not found")}
	UserNotFound              = Error{Err: errors.New("user not found")}
	NotEnoughCoins            = Error{Err: errors.New("not enough coins")}
	ReceiverNotFound          = Error{Err: errors.New("receiver not found")}
	CannotSentCoinsToYourself = Error{Err: errors.New("cannot send coins to yourself")}
)
