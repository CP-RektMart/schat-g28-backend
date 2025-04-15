package chat

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/chat"
	"github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/schat-g28-backend/internal/store"
)

type Handler struct {
	store          *store.Store
	chatService    *chat.Server
	authentication authentication.AuthMiddleware
}

func NewHandler(store *store.Store, authentication authentication.AuthMiddleware, chatService *chat.Server) *Handler {
	return &Handler{
		store:          store,
		chatService:    chatService,
		authentication: authentication,
	}
}
