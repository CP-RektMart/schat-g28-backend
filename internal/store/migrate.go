package store

import (
	"log/slog"

	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/pkg/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	logger.Info("Running migrations...")

	if err := db.AutoMigrate(
		&model.User{},
		&model.Group{},
		&model.DirectMessage{},
		&model.GroupMessage{},
		&model.File{},
	); err != nil {
		logger.Panic("failed to migrate store", slog.Any("error", err))
	}

	logger.Info("Migrations complete!")
}
