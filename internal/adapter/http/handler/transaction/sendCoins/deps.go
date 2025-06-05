package sendCoins

import (
	"context"
	"github.com/ALexfonSchneider/avito-shop/internal/dto"
)

type Service interface {
	SendCoin(ctx context.Context, request dto.SendCoinRequest) error
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{service}
}
