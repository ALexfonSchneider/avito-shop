package pgxpool

import (
	"context"
	"github.com/ALexfonSchneider/avito-shop/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func MustPGXPool(ctx context.Context, cfg *config.Config) *pgxpool.Pool {
	conf, err := pgxpool.ParseConfig(cfg.Postgres.ConnectionString())
	if err != nil {
		panic(err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		panic(err)
	}

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	return pool
}
