package buy

import (
	"github.com/ALexfonSchneider/avito-shop/internal/adapter/http/handler"
	"github.com/ALexfonSchneider/avito-shop/internal/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Request struct {
	Item string `param:"item"`
}

func (h *Handler) Handle(c echo.Context) error {
	var req Request
	if err := c.Bind(&req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	claims := handler.GetClaims(c)
	if claims == nil {
		return handler.Unauthorized
	}

	if err := h.service.BuyMerch(ctx, dto.BuyMerchRequest{
		MerchName: req.Item,
		UserID:    claims.UserID,
		Quantity:  1,
	}); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
