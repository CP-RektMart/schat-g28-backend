package redis

import (
	"context"
	"log/slog"

	"github.com/CP-RektMart/schat-g28-backend/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	URL string `env:"URL"`
}

func New(ctx context.Context, config Config) *redis.Client {
	opt, err := redis.ParseURL(config.URL)
	if err != nil {
		logger.Panic("failed to parse redis url", slog.Any("error", err))
	}
	rdb := redis.NewClient(opt)

	return rdb
}
