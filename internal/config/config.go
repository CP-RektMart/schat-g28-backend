package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"

	"github.com/CP-RektMart/schat-g28-backend/internal/jwt"
	"github.com/CP-RektMart/schat-g28-backend/internal/oauth"
	"github.com/CP-RektMart/schat-g28-backend/internal/server"
	"github.com/CP-RektMart/schat-g28-backend/pkg/logger"
	"github.com/CP-RektMart/schat-g28-backend/pkg/postgres"
	"github.com/CP-RektMart/schat-g28-backend/pkg/redis"
	"github.com/CP-RektMart/schat-g28-backend/pkg/storage"
)

type AppConfig struct {
	Server      server.Config      `envPrefix:"SERVER_"`
	Cors        server.CorsConfig  `envPrefix:"CORS_"`
	Logger      logger.Config      `envPrefix:"LOGGER_"`
	Postgres    postgres.Config    `envPrefix:"POSTGRES_"`
	Redis       redis.Config       `envPrefix:"REDIS_"`
	JWT         jwt.Config         `envPrefix:"JWT_"`
	OAuthGoogle oauth.GoogleConfig `envPrefix:"OAUTH_GOOGLE"`
	Store       storage.Config     `envPrefix:"STORAGE_"`
}

func Load() *AppConfig {
	appConfig := &AppConfig{}
	_ = godotenv.Load()

	if err := env.Parse(appConfig); err != nil {
		log.Panic("failed parse env", err)
	}

	return appConfig
}
