package infrastructure

import (
	"errors"
	"fmt"
)

type Error struct {
	Err error
}

func (e Error) Error() string {
	return fmt.Sprintf("infrastructure error: %v", e.Err)
}

func (e Error) Unwrap() error {
	return e.Err
}

var (
	Internal = Error{Err: errors.New("internal")}
)
