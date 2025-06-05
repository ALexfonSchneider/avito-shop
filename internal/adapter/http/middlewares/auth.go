package middlewares

import (
	authdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/auth"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	signingKey string
	algorithm  string
}

func NewAuthMiddleware(signingKey, algorithm string) *AuthMiddleware {
	return &AuthMiddleware{signingKey: signingKey, algorithm: algorithm}
}

func (m *AuthMiddleware) Middleware() echo.MiddlewareFunc {
	config := echojwt.Config{
		TokenLookup:   "header:Authorization",
		SigningMethod: m.algorithm,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(authdomain.Claims)
		},
		SigningKey: []byte(m.signingKey),
	}

	return echojwt.WithConfig(config)
}
