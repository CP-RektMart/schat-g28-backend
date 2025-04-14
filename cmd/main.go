package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/CP-RektMart/schat-g28-backend/internal/chat"
	"github.com/CP-RektMart/schat-g28-backend/internal/config"
	"github.com/CP-RektMart/schat-g28-backend/internal/database"
	"github.com/CP-RektMart/schat-g28-backend/internal/jwt"
	"github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/schat-g28-backend/internal/oauth"
	"github.com/CP-RektMart/schat-g28-backend/internal/server"
	"github.com/CP-RektMart/schat-g28-backend/internal/services/auth"
	"github.com/CP-RektMart/schat-g28-backend/internal/services/file"
	"github.com/CP-RektMart/schat-g28-backend/internal/services/message"
	"github.com/CP-RektMart/schat-g28-backend/pkg/logger"
	"github.com/CP-RektMart/schat-g28-backend/pkg/storage"
	"github.com/go-playground/validator/v10"
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
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	config, err := config.Load()
	if err != nil {
		logger.PanicContext(ctx, "failed to load config", "error", err)
	}

	if err := logger.Init(config.Logger); err != nil {
		logger.PanicContext(ctx, "failed to initialize logger", "error", err)
	}

	storage := storage.New(ctx, config.Store)
	db := database.NewDB(ctx, config.Postgres)
	store := database.New(ctx, config.Postgres, config.Redis)
	server := server.New(config.Server, config.Cors, config.JWT, store)
	validate := validator.New()

	// repository
	authRepo := auth.NewRepository(db)
	fileRepo := file.NewRepository(db)

	// services
	jwtService := jwt.New(config.JWT, store.Cache)
	chatService := chat.NewServer(store, validate)
	googleOauth := oauth.NewGoogle(config.OAuthGoogle)

	// middlewares
	authMiddleware := authentication.NewAuthMiddleware(jwtService)

	// handlers
	authHandler := auth.NewHandler(validate, authRepo, jwtService, authMiddleware, googleOauth)
	messageHandler := message.NewHandler(store, authMiddleware, chatService)
	fileHandler := file.NewHandler(storage, authMiddleware, fileRepo)
	server.RegisterDocs()

	// routes
	server.RegisterRoutes(
		authMiddleware,
		authHandler,
		messageHandler,
		fileHandler,
	)

	server.Start(ctx, stop)
}
