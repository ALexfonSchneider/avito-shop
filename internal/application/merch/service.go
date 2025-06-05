package merch

import (
	"context"
	"github.com/ALexfonSchneider/avito-shop/internal/application"
	merchdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/merch"
	"github.com/ALexfonSchneider/avito-shop/internal/dto"
	"time"
)

//go:generate mockgen -source=deps.go -destination deps_mock_test.go -package "${GOPACKAGE}_test"

func (s *Service) BuyMerch(ctx context.Context, request dto.BuyMerchRequest) error {
	merch, err := s.merch.FindMerchByName(ctx, request.MerchName)
	if err != nil {
		return err
	}

	if merch == nil {
		return application.MerchNotFound
	}

	user, err := s.user.FindUserByID(ctx, request.UserID)
	if err != nil {
		return err
	}

	if user == nil {
		return application.UserNotFound
	}

	purchase := merchdomain.NewPurchaseFromMerchAndUser(user.ID, *merch, request.Quantity, time.Now())
	if err := purchase.Validate(); err != nil {
		return err
	}

	if !user.CanBuy(purchase.Amount) {
		return application.NotEnoughCoins
	}

	if err := s.user.IncrementUserBalance(ctx, user.ID, -purchase.Amount); err != nil {
		return err
	}

	if err := s.merch.CreatePurchase(ctx, purchase); err != nil {
		return err
	}

	return nil
}
