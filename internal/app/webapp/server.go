package router

import (
	"context"
	"errors"
	"fmt"
	authhandler "github.com/ALexfonSchneider/avito-shop/internal/adapter/http/handler/auth/authenticate"
	buymerchhandler "github.com/ALexfonSchneider/avito-shop/internal/adapter/http/handler/merch/buy"
	sendcoinshandler "github.com/ALexfonSchneider/avito-shop/internal/adapter/http/handler/transaction/sendCoins"
	userinfohandler "github.com/ALexfonSchneider/avito-shop/internal/adapter/http/handler/user/userInfo"
	mymiddleware "github.com/ALexfonSchneider/avito-shop/internal/adapter/http/middlewares"
	authservice "github.com/ALexfonSchneider/avito-shop/internal/application/auth"
	merchservice "github.com/ALexfonSchneider/avito-shop/internal/application/merch"
	transactionservice "github.com/ALexfonSchneider/avito-shop/internal/application/transaction"
	userservice "github.com/ALexfonSchneider/avito-shop/internal/application/user"
	"github.com/ALexfonSchneider/avito-shop/internal/config"
	postgresrepo "github.com/ALexfonSchneider/avito-shop/internal/infrastructure/persistance/postgres"
	hasherrepo "github.com/ALexfonSchneider/avito-shop/internal/pkg/auth/hasher"
	jwtrepo "github.com/ALexfonSchneider/avito-shop/internal/pkg/auth/jwt"
	mypool "github.com/ALexfonSchneider/avito-shop/internal/pkg/pgxpool"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net"
	"net/http"
	"time"
)

type WebApp struct {
	pool   *pgxpool.Pool
	server *http.Server
	logger *slog.Logger

	done chan interface{}
}

func New(ctx context.Context, cfg *config.Config, log *slog.Logger) *WebApp {
	done := make(chan interface{})

	pool := mypool.MustPGXPool(ctx, cfg)

	postgresRepository := postgresrepo.New(pool)

	jwt := jwtrepo.NewTokenProvider(jwtrepo.Config{
		SecretKey:      cfg.Auth.GetSecretKey(),
		Issuer:         cfg.Auth.GetIssuer(),
		SigningMethod:  cfg.Auth.GetSigningMethod(),
		AccessTokenTTL: cfg.Auth.GetAccessTokenTTL() * time.Hour,
	})
	hasher := hasherrepo.New()

	authService := authservice.New(postgresRepository, jwt, hasher, log)
	merchService := merchservice.New(postgresRepository, postgresRepository)
	transactionService := transactionservice.New(postgresRepository, postgresRepository)
	userService := userservice.New(postgresRepository, postgresRepository)

	authHandler := authhandler.New(authService, log)
	buyMerchHandler := buymerchhandler.New(merchService, log)
	sendCoinsHandler := sendcoinshandler.New(transactionService)
	userInfoHandler := userinfohandler.New(userService)

	r := echo.New()

	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
	}))

	r.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				log.LogAttrs(ctx, slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			}
			return nil
		},
	}))
	r.Use(middleware.Recover())

	echoErrorHandler := mymiddleware.NewErrorsMiddleware()
	r.HTTPErrorHandler = echoErrorHandler.ErrorHandler // global error handler

	authMiddleware := mymiddleware.NewAuthMiddleware(cfg.Auth.GetSecretKey(), cfg.Auth.GetSigningMethod()).Middleware()

	api := r.Group("/api")

	{
		api.POST("/auth", authHandler.Handle)
	}

	api.GET("/buy/:item", buyMerchHandler.Handle, authMiddleware)
	api.POST("/sendCoin", sendCoinsHandler.Handle, authMiddleware)
	api.GET("/info", userInfoHandler.Handle, authMiddleware)

	r.Server.BaseContext = func(_ net.Listener) context.Context { return ctx }
	r.Server.Addr = fmt.Sprintf(":%d", cfg.App.Port)
	r.Server.DisableGeneralOptionsHandler = true
	r.Server.Handler = r

	return &WebApp{pool, r.Server, log, done}
}

func (app *WebApp) Start() error {
	var serverErr error

	if err := app.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		serverErr = err
	}

	app.done <- serverErr

	return serverErr
}

func (app *WebApp) Shutdown(ctx context.Context) error {
	defer app.pool.Close()

	app.logger.Info("http server shutting down")
	if err := app.server.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		app.logger.Error("http server shutdown failed")
	} else {
		return err
	}

	return nil
}

func (app *WebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.server.Handler.ServeHTTP(w, r)
}

func (app *WebApp) Done() <-chan interface{} {
	return app.done
}
