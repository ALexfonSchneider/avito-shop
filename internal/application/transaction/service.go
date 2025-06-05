package transaction

import (
	"context"
	"github.com/ALexfonSchneider/avito-shop/internal/application"
	transactiondomain "github.com/ALexfonSchneider/avito-shop/internal/domain/transaction"
	"github.com/ALexfonSchneider/avito-shop/internal/dto"
)

//go:generate mockgen -source=deps.go -destination deps_mock_test.go -package "${GOPACKAGE}_test"

func (s *Service) SendCoin(ctx context.Context, request dto.SendCoinRequest) error {
	userFrom, err := s.user.FindUserByID(ctx, request.FromUserID)
	if err != nil {
		return err
	}

	if userFrom == nil {
		return application.UserNotFound
	}

	if !userFrom.CanSend(request.Amount) {
		return application.NotEnoughCoins
	}

	userTo, err := s.user.FindUserByUsername(ctx, request.ToUserName)
	if err != nil {
		return err
	}

	if userTo == nil {
		return application.ReceiverNotFound
	}

	if userFrom.Username == userTo.Username {
		return application.CannotSentCoinsToYourself
	}

	if err = s.user.IncrementUserBalance(ctx, userFrom.ID, -request.Amount); err != nil {
		return err
	}
	if err = s.user.IncrementUserBalance(ctx, userTo.ID, request.Amount); err != nil {
		return err
	}

	transaction := transactiondomain.NewTransaction(userFrom.ID, userTo.ID, request.Amount)
	if err = s.transaction.CreateTransaction(ctx, transaction); err != nil {
		return err
	}

	return nil
}
