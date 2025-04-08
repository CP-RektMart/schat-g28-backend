package config

import (
	"log"

	"github.com/CP-RektMart/computer-network-g28/backend/internal/server"
	"github.com/CP-RektMart/computer-network-g28/backend/package/postgres"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Postgres postgres.Config `envPrefix:"POSTGRES_"`
	Server   server.Config   `envPrefix:"SERVER_"`
}

func Load() *AppConfig {
	appConfig := &AppConfig{}
	_ = godotenv.Load()

	if err := env.Parse(appConfig); err != nil {
		log.Fatalf("failed parse env: %s", err)
	}

	return appConfig
}
