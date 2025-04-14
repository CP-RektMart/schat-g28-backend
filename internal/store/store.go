package store

import (
	"context"
	"log/slog"

	"github.com/CP-RektMart/schat-g28-backend/pkg/logger"
	pglib "github.com/CP-RektMart/schat-g28-backend/pkg/postgres"
	rdlib "github.com/CP-RektMart/schat-g28-backend/pkg/redis"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func New(ctx context.Context, pgConfig pglib.Config, rdConfig rdlib.Config) *Store {
	db, err := gorm.Open(postgres.Open(pgConfig.String()), &gorm.Config{})
	if err != nil {
		logger.PanicContext(ctx, "failed to connect to store", slog.Any("error", err))
	}

	redisConn := rdlib.New(ctx, rdConfig)

	store := &Store{
		DB:    db,
		Cache: redisConn,
	}
	return store
}
