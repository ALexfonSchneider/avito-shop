package merch

import (
	"context"
	merchdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/merch"
	userdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/user"
)

type UserRepository interface {
	FindUserByID(ctx context.Context, ID string) (*userdomain.User, error)
	IncrementUserBalance(ctx context.Context, ID string, points int64) error
}

type MerchRepository interface {
	FindMerchByName(ctx context.Context, name string) (*merchdomain.Merch, error)
	CreatePurchase(ctx context.Context, merch *merchdomain.Purchase) error
}

type Service struct {
	user  UserRepository
	merch MerchRepository
}

func New(user UserRepository, merch MerchRepository) *Service {
	return &Service{
		user:  user,
		merch: merch,
	}
}
