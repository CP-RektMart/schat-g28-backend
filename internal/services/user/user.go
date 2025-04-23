package user

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/chat"
	"github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"
	"gorm.io/gorm"
)

type Handler struct {
	db             *gorm.DB
	authMiddleware authentication.AuthMiddleware
	repo           *Repository
	chatService    *chat.Server
}

func NewHandler(db *gorm.DB, authMiddleware authentication.AuthMiddleware, repo *Repository, chatService *chat.Server) *Handler {
	return &Handler{
		db:             db,
		authMiddleware: authMiddleware,
		repo:           repo,
		chatService:    chatService,
	}
}
