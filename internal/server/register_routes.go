package server

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/schat-g28-backend/internal/services/auth"
	"github.com/CP-RektMart/schat-g28-backend/internal/services/chat"
	"github.com/CP-RektMart/schat-g28-backend/internal/services/file"
	"github.com/CP-RektMart/schat-g28-backend/internal/services/group"
	"github.com/CP-RektMart/schat-g28-backend/internal/services/user"
	"github.com/gofiber/contrib/websocket"
)

func (s *Server) RegisterRoutes(
	authMiddleware authentication.AuthMiddleware,
	authHandler *auth.Handler,
	chatHandler *chat.Handler,
	fileHandler *file.Handler,
	groupHandler *group.Handler,
	userHandler *user.Handler,
) {
	v1 := s.app.Group("/api/v1")

	// auth
	auth := v1.Group("/auth")
	auth.Post("/login", authHandler.HandleLogin)
	auth.Post("/refresh-token", authHandler.HandleRefreshToken)
	auth.Post("/logout", authMiddleware.Auth, authHandler.HandleLogout)

	// me
	me := v1.Group("/me")
	me.Get("/", authMiddleware.Auth, authHandler.HandleGetMe)
	me.Patch("/", authMiddleware.Auth, authHandler.HandleUpdateMe)

	// messages
	chat := v1.Group("/chat")
	chat.Use("/ws", chatHandler.HandleSupportWebAPI, authMiddleware.Auth, chatHandler.HandleWebsocket)
	chat.Get("/ws", websocket.New(chatHandler.HandleRealTimeMessages))
	chat.Get("/", authMiddleware.Auth, chatHandler.HandleListChat)

	// file
	file := v1.Group("/files")
	file.Post("/", authMiddleware.Auth, fileHandler.HandleUploadFile)
	file.Delete("/:id", authMiddleware.Auth, fileHandler.HandleDeleteFile)

	// group
	group := v1.Group("/groups")
	group.Post("/", authMiddleware.Auth, groupHandler.HandleCreateGroup)
	group.Get("/", authMiddleware.Auth, groupHandler.HandleGetGroups)
	group.Get("/:id", authMiddleware.Auth, groupHandler.HandleGetByID)
	group.Patch("/:id", authMiddleware.Auth, groupHandler.HandleUpdateGroup)
	group.Get("/:id/join", authMiddleware.Auth, groupHandler.HandleJoinGroup)
	group.Get("/:id/leave", authMiddleware.Auth, groupHandler.HandleLeaveGroup)
	group.Delete("/:id", authMiddleware.Auth, groupHandler.HandleDeleteGroup)
	group.Post("/:groupID/members/:userID", authMiddleware.Auth, groupHandler.HandleAddGroupMember)
	group.Delete("/:groupID/members/:userID", authMiddleware.Auth, groupHandler.HandleKickGroupMember)

	// users
	users := v1.Group("/users")
	users.Get("/", authMiddleware.Auth, userHandler.HandleGetUsers)
	users.Get("/:id", authMiddleware.Auth, userHandler.HandleGetByID)
}
