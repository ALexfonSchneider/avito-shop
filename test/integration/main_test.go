//go:build integration
// +build integration

package integration

import (
	"context"
	webapp "github.com/ALexfonSchneider/avito-shop/internal/app/webapp"
	"github.com/ALexfonSchneider/avito-shop/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

var webApp *webapp.WebApp

func TestMain(m *testing.M) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.MustLoadTestConfig()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	webApp = webapp.New(ctx, cfg, logger)

	m.Run()
}
