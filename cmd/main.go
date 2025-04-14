package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/CP-RektMart/schat-g28-backend/internal/config"
	"github.com/CP-RektMart/schat-g28-backend/internal/jwt"
	"github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/schat-g28-backend/internal/oauth"
	"github.com/CP-RektMart/schat-g28-backend/internal/server"
	"github.com/CP-RektMart/schat-g28-backend/internal/services/auth"
	"github.com/CP-RektMart/schat-g28-backend/internal/services/file"
	"github.com/CP-RektMart/schat-g28-backend/internal/store"
	"github.com/CP-RektMart/schat-g28-backend/pkg/logger"
	"github.com/CP-RektMart/schat-g28-backend/pkg/redis"
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

	config := config.Load()

	logger.Init(config.Logger)

	storage := storage.New(ctx, config.Store)
	db := store.NewDB(ctx, config.Postgres)
	cache := redis.New(ctx, config.Redis)
	store.Migrate(db)

	server := server.New(config.Server, config.Cors, config.JWT)
	validate := validator.New()

	// repository
	authRepo := auth.NewRepository(db)
	fileRepo := file.NewRepository(db, storage)

	// services
	jwtService := jwt.New(config.JWT, cache)
	// chatService := chat.NewServer(store1, validate)
	googleOauth := oauth.NewGoogle(config.OAuthGoogle)

	// middlewares
	authMiddleware := authentication.NewAuthMiddleware(jwtService)

	// handlers
	authHandler := auth.NewHandler(validate, authRepo, jwtService, authMiddleware, googleOauth)
	// messageHandler := message.NewHandler(store1, authMiddleware, chatService)
	fileHandler := file.NewHandler(storage, authMiddleware, fileRepo)
	server.RegisterDocs()

	// routes
	server.RegisterRoutes(
		authMiddleware,
		authHandler,
		// messageHandler,
		fileHandler,
	)

	server.Start(ctx, stop)
}
