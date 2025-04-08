package database

import (
	"log"

	pglib "github.com/CP-RektMart/computer-network-g28/backend/package/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	PgConfig pglib.Config
}

type Store struct {
	Config Config
	DB     *gorm.DB
}

func New(config Config) *Store {
	DB, err := gorm.Open(postgres.Open(config.PgConfig.String()), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	return &Store{
		Config: config,
		DB:     DB,
	}
}
