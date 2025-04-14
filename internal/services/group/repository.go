package group

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(g model.Group) error {
	return r.db.Create(&g).Error
}
