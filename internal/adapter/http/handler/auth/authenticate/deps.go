package authenticate

import (
	"context"
	"github.com/ALexfonSchneider/avito-shop/internal/dto"
	"log/slog"
)

type Service interface {
	Authorize(ctx context.Context, request dto.AuthorizeRequest) (string, error)
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
