package user

import (
	"context"
	transactiondomain "github.com/ALexfonSchneider/avito-shop/internal/domain/transaction"
	userdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/user"
	"github.com/ALexfonSchneider/avito-shop/internal/dto"
)

type UserRepository interface {
	FindUserByID(ctx context.Context, ID string) (*userdomain.User, error)
	UserInventory(ctx context.Context, userID string) ([]dto.InventoryItem, error)
}

type TransactionRepository interface {
	GetTransactions(ctx context.Context, userID string) ([]transactiondomain.Transaction, error)
}

type Service struct {
	user        UserRepository
	transaction TransactionRepository
}

func New(user UserRepository, transaction TransactionRepository) *Service {
	return &Service{
		user:        user,
		transaction: transaction,
	}
}
