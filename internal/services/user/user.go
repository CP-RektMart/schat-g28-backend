package user

import (
	"errors"

	"github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"gorm.io/gorm"
)

type Handler struct {
	authMiddleware authentication.AuthMiddleware
	repo           *Repository
}

func NewHandler(authMiddleware authentication.AuthMiddleware, repo *Repository) *Handler {
	return &Handler{
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
