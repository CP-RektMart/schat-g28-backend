package user

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"
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
