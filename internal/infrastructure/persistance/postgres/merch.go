package postgres

import (
	"context"
	"errors"
	merchdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/merch"
	"github.com/ALexfonSchneider/avito-shop/internal/infrastructure"
)

func (r *Repository) CreateMerch(ctx context.Context, merch *merchdomain.Merch) error {
	const sql = "INSERT INTO merch (id, name, description, price, created_at) VALUES ($1, $2, $3, $4, $5)"

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return errors.Join(err, infrastructure.Internal)
	}

	if _, err = tx.Exec(ctx, sql, merch.Id, merch.Name, merch.Description, merch.Price, merch.CreatedAt); err != nil {
		return errors.Join(err, infrastructure.Internal)
	}

	if err = tx.Commit(ctx); err != nil {
		return errors.Join(err, infrastructure.Internal)
	}

	return nil
}

func (r *Repository) GetMerchByID(ctx context.Context, ID string) (*merchdomain.Merch, error) {
	const sql = "SELECT id, name, description, price, created_at FROM merch WHERE id=$1"

	rows, err := r.pool.Query(ctx, sql, ID)
	defer rows.Close()

	if err != nil {
		return nil, errors.Join(err, infrastructure.Internal)
	}

	if rows.Next() {
		var merch merchdomain.Merch
		err = rows.Scan(&merch.Id, &merch.Name, &merch.Description, &merch.Price, &merch.CreatedAt)
		if err != nil {
			return nil, errors.Join(err, infrastructure.Internal)
		}

		return &merch, nil
	}

	return nil, nil
}

func (r *Repository) FindMerchByName(ctx context.Context, name string) (*merchdomain.Merch, error) {
	const sql = "SELECT id, name, description, price, created_at FROM merch WHERE name=$1"

	rows, err := r.pool.Query(ctx, sql, name)
	defer rows.Close()

	if err != nil {
		return nil, errors.Join(err, infrastructure.Internal)
	}

	if rows.Next() {
		var merch merchdomain.Merch
		err = rows.Scan(&merch.Id, &merch.Name, &merch.Description, &merch.Price, &merch.CreatedAt)
		if err != nil {
			return nil, errors.Join(err, infrastructure.Internal)
		}

		return &merch, nil
	}

	return nil, nil
}

func (r *Repository) CreatePurchase(ctx context.Context, purchase *merchdomain.Purchase) error {
	const sql = "INSERT INTO purchases (id, user_id, merch_id, quantity, amount, purchased_at) VALUES ($1, $2, $3, $4, $5, $6)"

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return errors.Join(err, infrastructure.Internal)
	}

	if _, err := tx.Exec(ctx, sql, purchase.ID, purchase.UserID, purchase.MerchID, purchase.Quantity, purchase.Amount, purchase.PurchasedAt); err != nil {
		return errors.Join(err, infrastructure.Internal)
	}

	if err := tx.Commit(ctx); err != nil {
		return errors.Join(err, infrastructure.Internal)
	}

	return nil
}
