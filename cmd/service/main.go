package main

import (
	"context"
	webapp "github.com/ALexfonSchneider/avito-shop/internal/app/webapp"
	"github.com/ALexfonSchneider/avito-shop/internal/config"
	"github.com/ALexfonSchneider/avito-shop/internal/pkg/logger"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	cfg := config.MustConfig()

	log := logger.MustLogger(logger.Config{IncludeProgramInfo: true, Level: slog.LevelInfo})

	webApp := webapp.New(ctx, cfg, log)

	var err error

	go func() {
		log.Info("starting webapp")
		if err = webApp.Start(); err != nil {
			log.Error("Failed to start web app", "error", err)
		}
	}()

	select {
	case <-ctx.Done():
	case <-webApp.Done():
	}

	if err := webApp.Shutdown(ctx); err != nil {
		log.Error("Failed to shutdown web app", "error", err)
	}
}
