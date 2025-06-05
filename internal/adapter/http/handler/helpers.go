package handler

import (
	authdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetClaims(c echo.Context) *authdomain.Claims {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil
	}

	claims, ok := token.Claims.(*authdomain.Claims)
	if !ok {
		return nil
	}

	return claims
}
