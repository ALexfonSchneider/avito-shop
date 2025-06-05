package auth

import (
	"errors"
	"github.com/ALexfonSchneider/avito-shop/internal/domain"
)

var (
	UnexpectedSigningMethodError = domain.Error{Err: errors.New("unexpected signing method")}
)
