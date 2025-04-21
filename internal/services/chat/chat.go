package chat

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/chat"
	"github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"
	"gorm.io/gorm"
)

type Handler struct {
	db             *gorm.DB
	chatService    *chat.Server
	authMiddleware authentication.AuthMiddleware
}

func NewHandler(db *gorm.DB, authMiddleware authentication.AuthMiddleware, chatService *chat.Server) *Handler {
	return &Handler{
		db:             db,
		chatService:    chatService,
		authMiddleware: authMiddleware,
	}
}
