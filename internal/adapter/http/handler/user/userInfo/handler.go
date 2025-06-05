package userInfo

import (
	"github.com/ALexfonSchneider/avito-shop/internal/adapter/http/handler"
	"github.com/labstack/echo/v4"
	"net/http"
)

type InventoryItem struct {
	Type     string `json:"type"`
	Quantity int32  `json:"quantity"`
}

type CoinHistoryReceived struct {
	FromUser string `json:"fromUser"`
	Amount   int64  `json:"amount"`
}

type CoinHistorySent struct {
	ToUser string `json:"toUser"`
	Amount int64  `json:"amount"`
}

type CoinHistory struct {
	Received []CoinHistoryReceived `json:"received"`
	Sent     []CoinHistorySent     `json:"sent"`
}

type Response struct {
	Coins       int64           `json:"coins"`
	Inventory   []InventoryItem `json:"inventory"`
	CoinHistory CoinHistory     `json:"coinHistory"`
}

func (h *Handler) Handle(c echo.Context) error {
	claims := handler.GetClaims(c)
	if claims == nil {
		return handler.Unauthorized
	}

	ctx := c.Request().Context()
	info, err := h.service.GetUserInfo(ctx, claims.UserID)
	if err != nil {
		return err
	}

	resp := Response{
		Coins: info.Coins,
		Inventory: func() (items []InventoryItem) {
			for _, item := range info.Inventory {
				items = append(items, InventoryItem{
					Type:     item.Name,
					Quantity: item.Quantity,
				})
			}
			return
		}(),
		CoinHistory: CoinHistory{
			Received: func() (received []CoinHistoryReceived) {
				for _, record := range info.History.Received {
					received = append(received, CoinHistoryReceived{
						FromUser: record.FromUser,
						Amount:   record.Amount,
					})
				}
				return
			}(),
			Sent: func() (sent []CoinHistorySent) {
				for _, record := range info.History.Sent {
					sent = append(sent, CoinHistorySent{
						ToUser: record.ToUser,
						Amount: record.Amount,
					})
				}
				return
			}(),
		},
	}

	return c.JSON(http.StatusOK, resp)
}
