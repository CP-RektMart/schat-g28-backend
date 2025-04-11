package config

import (
	"log"

	"github.com/CP-RektMart/schat-g28-backend/internal/jwt"
	"github.com/CP-RektMart/schat-g28-backend/internal/server"
	"github.com/CP-RektMart/schat-g28-backend/pkg/logger"
	"github.com/CP-RektMart/schat-g28-backend/pkg/postgres"
	"github.com/CP-RektMart/schat-g28-backend/pkg/redis"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Server         server.Config     `envPrefix:"SERVER_"`
	Logger         logger.Config     `envPrefix:"LOGGER_"`
	Postgres       postgres.Config   `envPrefix:"POSTGRES_"`
	Redis          redis.Config      `envPrefix:"REDIS_"`
	Cors           server.CorsConfig `envPrefix:"CORS_"`
	JWT            jwt.Config        `envPrefix:"JWT_"`
	GoogleClientID string            `env:"GOOGLE_CLIENT_ID"`
	FrontendURL    string            `env:"FRONTEND_URL"`
}

func Load() *AppConfig {
	appConfig := &AppConfig{}
	_ = godotenv.Load()

	if err := env.Parse(appConfig); err != nil {
		log.Fatalf("failed parse env: %s", err)
	}

	return appConfig
}
