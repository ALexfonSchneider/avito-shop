package domain

import (
	"fmt"
)

type Error struct {
	Err error
}

func (e Error) Error() string {
	return fmt.Sprintf("domain error: %v", e.Err)
}
