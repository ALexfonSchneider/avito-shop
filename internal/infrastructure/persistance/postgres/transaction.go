package postgres

import (
	"context"
	"errors"
	transactiondomain "github.com/ALexfonSchneider/avito-shop/internal/domain/transaction"
	"github.com/ALexfonSchneider/avito-shop/internal/infrastructure"
)

func (r *Repository) CreateTransaction(ctx context.Context, transaction *transactiondomain.Transaction) error {
	const sql = "INSERT INTO transactions (id, sender_id, receiver_id, amount, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return errors.Join(err, infrastructure.Internal)
	}

	if err = tx.QueryRow(ctx, sql, transaction.ID, transaction.SenderID,
		transaction.ReceiverID, transaction.Amount, transaction.CreatedAt,
	).Scan(&transaction.ID); err != nil {
		return errors.Join(err, infrastructure.Internal)
	}

	if err = tx.Commit(ctx); err != nil {
		return errors.Join(err, infrastructure.Internal)
	}

	return nil
}

func (r *Repository) GetTransactions(ctx context.Context, userID string) ([]transactiondomain.Transaction, error) {
	const sql = "SELECT id, sender_id, receiver_id, amount, created_at FROM transactions WHERE receiver_id = $1 or sender_id = $1 order by created_at"

	rows, err := r.pool.Query(ctx, sql, userID)
	defer rows.Close()

	if err != nil {
		return nil, errors.Join(err, infrastructure.Internal)
	}

	var transactions []transactiondomain.Transaction
	for rows.Next() {
		var transaction transactiondomain.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.SenderID, &transaction.ReceiverID, &transaction.Amount, &transaction.CreatedAt); err != nil {
			return nil, errors.Join(err, infrastructure.Internal)
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
