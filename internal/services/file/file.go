package file

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/schat-g28-backend/pkg/storage"
)

type Handler struct {
	store          *storage.Client
	authMiddleware authentication.AuthMiddleware
	repo           *Repository
}

func NewHandler(
	store *storage.Client,
	authMiddleware authentication.AuthMiddleware,
	repo *Repository,
) *Handler {
	return &Handler{
		store:          store,
		authMiddleware: authMiddleware,
		repo:           repo,
	}
}
