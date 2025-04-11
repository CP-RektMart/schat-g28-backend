package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/CP-RektMart/schat-g28-backend/internal/chat"
	"github.com/CP-RektMart/schat-g28-backend/internal/config"
	"github.com/CP-RektMart/schat-g28-backend/internal/database"
	"github.com/CP-RektMart/schat-g28-backend/internal/jwt"
	"github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/schat-g28-backend/internal/server"
	"github.com/CP-RektMart/schat-g28-backend/internal/services/auth"
	"github.com/CP-RektMart/schat-g28-backend/internal/services/message"
	"github.com/CP-RektMart/schat-g28-backend/internal/services/user"
	"github.com/CP-RektMart/schat-g28-backend/internal/validator"
	"github.com/CP-RektMart/schat-g28-backend/pkg/logger"
)

// @title						Pic Me Pls API
// @version						0.1
// @description					Pic Me Pls API Documentation
// @securityDefinitions.apikey ApiKeyAuth
// @in							header
// @name						Authorization
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	// hello
	config := config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := logger.Init(config.Logger); err != nil {
		logger.PanicContext(ctx, "failed to initialize logger", slog.Any("error", err))
	}

	store := database.New(ctx, config.Postgres, config.Redis)
	server := server.New(config.Server, config.Cors, config.JWT, store)
	validate := validator.New()

	// services
	jwtService := jwt.New(config.JWT, store.Cache)
	chatService := chat.NewServer(store, validate)

	// middlewares
	authMiddleware := authentication.NewAuthMiddleware(jwtService)

	// handlers
	authHandler := auth.NewHandler(store, validate, jwtService, authMiddleware, config.GoogleClientID)
	userHandler := user.NewHandler(store, validate, authMiddleware)
	messageHandler := message.NewHandler(store, authMiddleware, chatService)
	server.RegisterDocs()

	// routes
	server.RegisterRoutes(
		authMiddleware,
		authHandler,
		userHandler,
		messageHandler,
	)

	server.Start(ctx, stop)
}
