package user

import (
	"context"
	"github.com/ALexfonSchneider/avito-shop/internal/application"
	"github.com/ALexfonSchneider/avito-shop/internal/dto"
)

//go:generate mockgen -source=deps.go -destination deps_mock_test.go -package "${GOPACKAGE}_test"

func (s *Service) GetUserInfo(ctx context.Context, userID string) (*dto.UserInfo, error) {
	user, err := s.user.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, application.UserNotFound
	}

	inventory, err := s.user.UserInventory(ctx, userID)
	if err != nil {
		return nil, err
	}

	transactions, err := s.transaction.GetTransactions(ctx, userID)
	if err != nil {
		return nil, err
	}

	var receivedTransactions []dto.CoinReceived
	var sentTransactions []dto.CoinSent
	for _, transaction := range transactions {
		if transaction.SenderID != userID {
			receivedTransactions = append(receivedTransactions, dto.CoinReceived{
				FromUser: transaction.SenderID,
				Amount:   transaction.Amount,
			})
		} else {
			sentTransactions = append(sentTransactions, dto.CoinSent{
				ToUser: transaction.ReceiverID,
				Amount: transaction.Amount,
			})
		}
	}

	return &dto.UserInfo{
		Coins:     user.Balance,
		Inventory: inventory,
		History: dto.CoinHistory{
			Received: receivedTransactions,
			Sent:     sentTransactions,
		},
	}, nil
}
