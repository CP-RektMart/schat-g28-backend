package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/cockroachdb/errors"
	"github.com/joho/godotenv"

	"github.com/CP-RektMart/schat-g28-backend/internal/jwt"
	"github.com/CP-RektMart/schat-g28-backend/internal/server"
	"github.com/CP-RektMart/schat-g28-backend/internal/utils/oauth"
	"github.com/CP-RektMart/schat-g28-backend/pkg/logger"
	"github.com/CP-RektMart/schat-g28-backend/pkg/postgres"
	"github.com/CP-RektMart/schat-g28-backend/pkg/redis"
)

type AppConfig struct {
	Server      server.Config      `envPrefix:"SERVER_"`
	Cors        server.CorsConfig  `envPrefix:"CORS_"`
	Logger      logger.Config      `envPrefix:"LOGGER_"`
	Postgres    postgres.Config    `envPrefix:"POSTGRES_"`
	Redis       redis.Config       `envPrefix:"REDIS_"`
	JWT         jwt.Config         `envPrefix:"JWT_"`
	OAuthGoogle oauth.GoogleConfig `envPrefix:"OAUTH_GOOGLE"`
}

func Load() (*AppConfig, error) {
	appConfig := &AppConfig{}
	_ = godotenv.Load()

	if err := env.Parse(appConfig); err != nil {
		return nil, errors.Wrap(err, "failed parse env")
	}

	return appConfig, nil
}
