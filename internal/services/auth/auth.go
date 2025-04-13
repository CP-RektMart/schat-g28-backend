package auth

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/database"
	"github.com/CP-RektMart/schat-g28-backend/internal/jwt"
	"github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/schat-g28-backend/internal/utils/oauth"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store          *database.Store
	validate       *validator.Validate
	jwtService     *jwt.JWT
	authmiddleware authentication.AuthMiddleware
	googleOauth    oauth.OAuth
}

func NewHandler(
	store *database.Store,
	validate *validator.Validate,
	jwtService *jwt.JWT,
	authmiddleware authentication.AuthMiddleware,
	googleOauth oauth.OAuth,
) *Handler {
	return &Handler{
		store:          store,
		validate:       validate,
		jwtService:     jwtService,
		authmiddleware: authmiddleware,
		googleOauth:    googleOauth,
	}
}
