package user

import (
	"errors"

	"github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/internal/utils/repository"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"gorm.io/gorm"
)

type Handler struct {
	db             *gorm.DB
	authMiddleware authentication.AuthMiddleware
	repo           *Repository
}

func NewHandler(db *gorm.DB, authMiddleware authentication.AuthMiddleware, repo *Repository) *Handler {
	return &Handler{
		db:             db,
		authMiddleware: authMiddleware,
		repo:           repo,
	}
}

func (r *Repository) Get() ([]model.User, error) {
	var u []model.User

	if err := r.db.Find(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.User{}, apperror.NotFound("groups not found", err)
		}
		return []model.User{}, err
	}

	return u, nil
}

func (r *Repository) GetByID(id uint, preload ...string) (model.User, error) {
	var u model.User

	db := repository.AccumulatePreload(r.db, preload...)

	if err := db.First(&u, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, apperror.NotFound("user not found", err)
		}
		return model.User{}, err
	}
	return u, nil
}
