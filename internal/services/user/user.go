package user

import (
	"github.com/CP-RektMart/computer-network-g28/backend/internal/database"
	"github.com/CP-RektMart/computer-network-g28/backend/internal/middlewares/authentication"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store          *database.Store
	validate       *validator.Validate
	authMiddleware authentication.AuthMiddleware
}

func NewHandler(store *database.Store, validate *validator.Validate, authMiddle authentication.AuthMiddleware) *Handler {
	return &Handler{
		store:          store,
		validate:       validate,
		authMiddleware: authMiddle,
	}
}
