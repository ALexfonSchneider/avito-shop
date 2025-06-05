package userInfo

import (
	"context"
	"github.com/ALexfonSchneider/avito-shop/internal/dto"
)

type Service interface {
	GetUserInfo(ctx context.Context, userID string) (*dto.UserInfo, error)
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{service}
}
