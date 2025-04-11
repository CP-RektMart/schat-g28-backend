package message

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/chat"
	"github.com/CP-RektMart/schat-g28-backend/internal/database"
	"github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"
)

type Handler struct {
	store          *database.Store
	chatService    *chat.Server
	authentication authentication.AuthMiddleware
}

func NewHandler(store *database.Store, authentication authentication.AuthMiddleware, chatService *chat.Server) *Handler {
	return &Handler{
		store:          store,
		chatService:    chatService,
		authentication: authentication,
	}
}
