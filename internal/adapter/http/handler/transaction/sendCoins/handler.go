package sendCoins

import (
	"github.com/ALexfonSchneider/avito-shop/internal/adapter/http/handler"
	"github.com/ALexfonSchneider/avito-shop/internal/dto"
	"github.com/labstack/echo/v4"
)

type Request struct {
	ToUser string `json:"toUser"`
	Amount int64  `json:"amount"`
}

func (h *Handler) Handle(c echo.Context) error {
	var req Request
	if err := c.Bind(&req); err != nil {
		return err
	}

	claims := handler.GetClaims(c)
	if claims == nil {
		return handler.Unauthorized
	}

	ctx := c.Request().Context()
	if err := h.service.SendCoin(ctx, dto.SendCoinRequest{
		FromUserID: claims.UserID,
		ToUserName: req.ToUser,
		Amount:     req.Amount,
	}); err != nil {
		return err
	}

	return nil
}
