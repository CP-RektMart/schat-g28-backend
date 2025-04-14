package group

import "github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"

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
