package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserCredentials struct {
	UserID string
}

func NewUserCredentials(userID string) *UserCredentials {
	return &UserCredentials{UserID: userID}
}

type TokenType string

const (
	AccessToken TokenType = "access_token"
)

type Claims struct {
	jwt.RegisteredClaims
	Use TokenType
	UserCredentials
}
