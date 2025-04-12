package file

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/database"
	"github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"
)

type Handler struct {
	store          *database.Store
	authMiddleware authentication.AuthMiddleware
}

func NewHandler(
	store *database.Store,
	authMiddleware authentication.AuthMiddleware,
) *Handler {
	return &Handler{
		store:          store,
		authMiddleware: authMiddleware,
	}
}
