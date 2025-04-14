package friend

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/schat-g28-backend/internal/services/auth"
)

type Handler struct {
	authMiddleware authentication.AuthMiddleware
	authRepo       *auth.Repository
}

func NewHandler(
	authMiddleware authentication.AuthMiddleware,
	authRepo *auth.Repository,
) *Handler {
	return &Handler{
		authMiddleware: authMiddleware,
		authRepo:       authRepo,
	}
}
