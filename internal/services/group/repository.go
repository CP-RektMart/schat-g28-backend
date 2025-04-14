package group

import (
	"errors"

	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
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

func (r *Repository) Get(id uint) (model.Group, error) {
	var g model.Group
	if err := r.db.Preload("Owner").Preload("Messages").Preload("Members").First(&g, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Group{}, apperror.NotFound("group not found", err)
		}
		return model.Group{}, err
	}
	return g, nil
}
