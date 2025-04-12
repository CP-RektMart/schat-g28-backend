package server

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/schat-g28-backend/internal/services/auth"
	"github.com/CP-RektMart/schat-g28-backend/internal/services/message"
	"github.com/CP-RektMart/schat-g28-backend/internal/services/user"
	"github.com/gofiber/contrib/websocket"
)

func (s *Server) RegisterRoutes(
	authMiddleware authentication.AuthMiddleware,
	authHandler *auth.Handler,
	userHandler *user.Handler,
	messageHandler *message.Handler,
) {
	v1 := s.app.Group("/api/v1")

	// auth
	auth := v1.Group("/auth")
	auth.Post("/login", authHandler.HandleLogin)
	auth.Post("/refresh-token", authHandler.HandleRefreshToken)
	auth.Post("/logout", authMiddleware.Auth, authHandler.HandleLogout)

	// all
	{
		all := v1.Group("/")

		// me
		me := all.Group("/me")
		me.Get("/", authMiddleware.Auth, userHandler.HandleGetMe)

		// messages
		message := all.Group("/messages")
		message.Use("/ws", messageHandler.HandleSupportWebAPI, authMiddleware.Auth, messageHandler.HandleWebsocket)
		message.Get("/ws", websocket.New(messageHandler.HandleRealTimeMessages))
		message.Get("/", authMiddleware.Auth, messageHandler.HandleListMessages)
	}
}
