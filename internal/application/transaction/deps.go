package transaction

import (
	"context"
	transactiondomain "github.com/ALexfonSchneider/avito-shop/internal/domain/transaction"
	userdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/user"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *transactiondomain.Transaction) error
}

type UserRepository interface {
	IncrementUserBalance(ctx context.Context, ID string, points int64) error
	FindUserByID(ctx context.Context, name string) (*userdomain.User, error)
	FindUserByUsername(ctx context.Context, name string) (*userdomain.User, error)
}

type Service struct {
	user        UserRepository
	transaction TransactionRepository
}

func New(transaction TransactionRepository, user UserRepository) *Service {
	return &Service{transaction: transaction, user: user}
}
