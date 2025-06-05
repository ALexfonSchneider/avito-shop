package main

import (
	"context"
	"github.com/ALexfonSchneider/avito-shop/cmd/init/merch"
	"github.com/ALexfonSchneider/avito-shop/internal/config"
	postgresrepo "github.com/ALexfonSchneider/avito-shop/internal/infrastructure/persistance/postgres"
	"github.com/ALexfonSchneider/avito-shop/internal/pkg/pgxpool"
)

func main() {
	ctx := context.Background()

	cfg := config.MustConfig()

	pool := pgxpool.MustPGXPool(ctx, cfg)
	defer pool.Close()

	postgresRepository := postgresrepo.New(pool)

	initMerch := merch.New(postgresRepository)

	if err := initMerch.Init(ctx); err != nil {
		panic(err)
	}
}
