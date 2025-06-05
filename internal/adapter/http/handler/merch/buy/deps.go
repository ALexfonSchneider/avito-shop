package buy

import (
	"context"
	"github.com/ALexfonSchneider/avito-shop/internal/dto"
	"log/slog"
)

type Service interface {
	BuyMerch(ctx context.Context, request dto.BuyMerchRequest) error
}

type Handler struct {
	service Service

	logger *slog.Logger
}

func New(service Service, logger *slog.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}
