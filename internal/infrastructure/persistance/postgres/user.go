package postgres

import (
	"context"
	"errors"
	userdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/user"
	"github.com/ALexfonSchneider/avito-shop/internal/infrastructure"
)

func (r *Repository) FindUserByID(ctx context.Context, ID string) (*userdomain.User, error) {
	const sql = "SELECT id, username, password_hash, balance, created_at FROM users WHERE id = $1"

	rows, err := r.pool.Query(ctx, sql, ID)
	defer rows.Close()

	if err != nil {
		return nil, errors.Join(err, infrastructure.Internal)
	}

	if rows.Next() {
		var user userdomain.User
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Balance, &user.CreatedAt)
		if err != nil {
			return nil, errors.Join(err, infrastructure.Internal)
		}

		return &user, nil
	}

	return nil, nil
}

func (r *Repository) FindUserByUsername(ctx context.Context, username string) (*userdomain.User, error) {
	const sql = "SELECT id, username, password_hash, balance, created_at FROM users WHERE username = $1"

	rows, err := r.pool.Query(ctx, sql, username)
	defer rows.Close()

	if err != nil {
		return nil, errors.Join(err, infrastructure.Internal)
	}

	if rows.Next() {
		var user userdomain.User
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Balance, &user.CreatedAt)
		if err != nil {
			return nil, errors.Join(err, infrastructure.Internal)
		}

		return &user, nil
	}

	return nil, nil
}

func (r *Repository) FindUsersByIDs(ctx context.Context, IDs []string) ([]userdomain.User, error) {
	const sql = "SELECT id, username, password_hash, balance, created_at FROM users WHERE id IN ($1)"

	rows, err := r.pool.Query(ctx, sql, IDs)
	defer rows.Close()

	if err != nil {
		return nil, errors.Join(err, infrastructure.Internal)
	}

	users := make([]userdomain.User, 0, len(IDs))
	for rows.Next() {
		var user userdomain.User
		if err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Balance, &user.CreatedAt); err != nil {
			return nil, errors.Join(err, infrastructure.Internal)
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *Repository) CreateUser(ctx context.Context, user *userdomain.User) error {
	const sql = "INSERT INTO users (id, username, password_hash, balance, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return errors.Join(err, infrastructure.Internal)
	}

	if err := tx.QueryRow(
		ctx, sql, user.ID, user.Username, user.Password,
		user.Balance, user.CreatedAt,
	).Scan(&user.ID); err != nil {
		return errors.Join(err, infrastructure.Internal)
	}

	if err := tx.Commit(ctx); err != nil {
		return errors.Join(err, infrastructure.Internal)
	}

	return nil
}

func (r *Repository) IncrementUserBalance(ctx context.Context, ID string, points int64) error {
	const sql = "UPDATE users SET balance = balance + $1 WHERE id = $2"

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return errors.Join(err, infrastructure.Internal)
	}

	if _, err := tx.Exec(ctx, sql, points, ID); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return errors.Join(err, infrastructure.Internal)
	}

	return nil
}
