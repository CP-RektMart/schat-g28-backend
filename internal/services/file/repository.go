package file

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

func (r *Repository) Create(f model.File) (model.File, error) {
	if err := r.db.Create(&f).Error; err != nil {
		return model.File{}, err
	}
	return f, nil
}
