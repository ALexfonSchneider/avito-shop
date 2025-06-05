package auth

import (
	"context"
	authdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/auth"
	userdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/user"
	"github.com/ALexfonSchneider/avito-shop/internal/dto"
	"time"
)

//go:generate mockgen -source=deps.go -destination deps_mock_test.go -package "${GOPACKAGE}_test"

// Authorize creating token for user. Create user if not exists.
func (s *Service) Authorize(ctx context.Context, request dto.AuthorizeRequest) (string, error) {
	user, err := s.user.FindUserByUsername(ctx, request.Username)
	if err != nil {
		return "", err
	}

	if user == nil {
		hashedPassword, err := s.hash.Hash(request.Password)
		if err != nil {
			return "", err
		}

		user = userdomain.NewUser(request.Username, hashedPassword, 1000, time.Now())
		if err = s.user.CreateUser(ctx, user); err != nil {
			return "", err
		}
	}

	if err = s.hash.Compare(request.Password, user.Password); err != nil {
		return "", err
	}

	claims := authdomain.NewUserCredentials(user.ID)
	token, err := s.jwt.CreateToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}
