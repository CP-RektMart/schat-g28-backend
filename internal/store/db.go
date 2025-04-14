package store

import (
	"context"
	"log/slog"

	"github.com/CP-RektMart/schat-g28-backend/pkg/logger"
	pglib "github.com/CP-RektMart/schat-g28-backend/pkg/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(ctx context.Context, pgConfig pglib.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(pgConfig.String()), &gorm.Config{})
	if err != nil {
		logger.PanicContext(ctx, "failed to connect to store", slog.Any("error", err))
	}
	return db
}
