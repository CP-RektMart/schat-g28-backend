package auth

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/jwt"
	"github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/schat-g28-backend/internal/oauth"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	repo           *Repository
	validate       *validator.Validate
	jwtService     *jwt.JWT
	authMiddleware authentication.AuthMiddleware
	googleOauth    oauth.OAuth
}

func NewHandler(
	validate *validator.Validate,
	repo *Repository,
	jwtService *jwt.JWT,
	authmiddleware authentication.AuthMiddleware,
	googleOauth oauth.OAuth,
) *Handler {
	return &Handler{
		validate:       validate,
		repo:           repo,
		jwtService:     jwtService,
		authMiddleware: authmiddleware,
		googleOauth:    googleOauth,
	}
}
