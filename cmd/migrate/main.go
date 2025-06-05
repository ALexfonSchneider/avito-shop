package main

import (
	"context"
	"github.com/ALexfonSchneider/avito-shop/internal/config"
	"github.com/ALexfonSchneider/avito-shop/internal/infrastructure/persistance/postgres"
	mypool "github.com/ALexfonSchneider/avito-shop/internal/pkg/pgxpool"
)

func main() {
	ctx := context.Background()
	cfg := config.MustConfig()

	pool := mypool.MustPGXPool(ctx, cfg)

	repo := postgres.New(pool)

	err := repo.Migrate(cfg)
	if err != nil {
		panic(err)
	}
}
