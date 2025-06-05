package jwt

import (
	authdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

type Config struct {
	SecretKey      string
	Issuer         string
	SigningMethod  string
	AccessTokenTTL time.Duration
}

type TokenProvider struct {
	cfg Config
}

func NewTokenProvider(config Config) *TokenProvider {
	return &TokenProvider{cfg: config}
}

func (t *TokenProvider) CreateToken(user *authdomain.UserCredentials) (string, error) {
	if user == nil {
		return "", errors.New("user is nil")
	}

	claims := authdomain.Claims{
		Use: authdomain.AccessToken,
		UserCredentials: authdomain.UserCredentials{
			UserID: user.UserID,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    t.cfg.Issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.cfg.AccessTokenTTL)),
			ID:        uuid.NewString(),
			Subject:   user.UserID,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(t.cfg.SigningMethod), claims)

	return token.SignedString([]byte(t.cfg.SecretKey))
}

func (t *TokenProvider) ValidateToken(tokenString string) (*authdomain.Claims, error) {
	claims := authdomain.Claims{}

	if _, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.GetSigningMethod(t.cfg.SigningMethod) {
			return nil, errors.Wrap(
				authdomain.UnexpectedSigningMethodError,
				token.Method.Alg(),
			)
		}
		return []byte(t.cfg.SecretKey), nil
	}); err != nil {
		return nil, err
	}

	return &claims, nil
}
