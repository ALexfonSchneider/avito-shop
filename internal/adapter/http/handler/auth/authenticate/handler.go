package authenticate

import (
	"github.com/ALexfonSchneider/avito-shop/internal/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Token string `json:"token"`
}

func (h *Handler) Handle(c echo.Context) error {
	var req Request
	if err := c.Bind(&req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	token, err := h.service.Authorize(ctx, dto.AuthorizeRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return err
	}

	if err = c.JSON(http.StatusOK, Response{Token: token}); err != nil {
		return err
	}

	return nil
}
