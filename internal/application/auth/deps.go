package auth

import (
	"context"
	authdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/auth"
	userdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/user"
	"log/slog"
)

type UserRepository interface {
	FindUserByUsername(ctx context.Context, username string) (*userdomain.User, error)
	CreateUser(ctx context.Context, user *userdomain.User) error
}

type JWTProvider interface {
	CreateToken(user *authdomain.UserCredentials) (string, error)
}

type Hasher interface {
	Hash(password string) (string, error)
	Compare(password, hash string) error
}

type Service struct {
	user UserRepository
	jwt  JWTProvider
	hash Hasher

	logger *slog.Logger
}

func New(user UserRepository, jwt JWTProvider, hash Hasher, logger *slog.Logger) *Service {
	subLog := logger.With("layer", "authenticate")

	return &Service{user: user, jwt: jwt, hash: hash, logger: subLog}
}
