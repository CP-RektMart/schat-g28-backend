package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	PostgresURL string
	RedisURL    string
}

type Store struct {
	Config Config
	DB     *gorm.DB
}

func New(config Config) *Store {
	DB, err := gorm.Open(postgres.Open(config.PostgresURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	return &Store{
		Config: config,
		DB:     DB,
	}
}
