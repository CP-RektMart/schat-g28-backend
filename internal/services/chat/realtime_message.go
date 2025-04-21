package chat

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"sync"

	"github.com/CP-RektMart/schat-g28-backend/internal/chat"
	"github.com/CP-RektMart/schat-g28-backend/internal/jwt"
	"github.com/CP-RektMart/schat-g28-backend/pkg/logger"
	"github.com/gofiber/contrib/websocket"
)

func (h *Handler) HandleRealTimeMessages(c *websocket.Conn) {
	jwtEntity, ok := c.Locals(jwtEntityKey).(jwt.JWTentity)
	if !ok {
		logger.Error("failed receive userID from jwtEntity")
		return
	}

	client := h.chatService.Register(jwtEntity.ID)

	var wg sync.WaitGroup
	wg.Add(2)

	go h.receiveRealtimeMessage(&wg, c, jwtEntity.ID)
	go h.sendRealtimeMessage(&wg, c, jwtEntity.ID, client)

	wg.Wait()
	c.Close()
}

func (h *Handler) receiveRealtimeMessage(wg *sync.WaitGroup, c *websocket.Conn, userID uint) {
	defer wg.Done()

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			logger.Error("failed receiving message", slog.Any("error", err))
			logger.Info("closing connection...")
			h.chatService.Logout(userID)
			break
		}

		if mt == websocket.TextMessage {
			// Determine if the message is for a group or direct chat
			var msgReq map[string]interface{}
			if err := json.Unmarshal(msg, &msgReq); err != nil {
				logger.Error("Failed to unmarshal message", slog.Any("error", err))
				continue
			}

			if _, ok := msgReq["groupId"]; ok {
				// Handle group message
				h.chatService.SendGroupRawString(userID, string(msg))
			} else {
				// Handle direct message
				h.chatService.SendDirectRawString(userID, string(msg))
			}
		}
	}
}

func (h *Handler) sendRealtimeMessage(wg *sync.WaitGroup, c *websocket.Conn, userID uint, client *chat.Client) {
	defer wg.Done()

	for {
		select {
		case <-client.Terminate:
			return
		case msg := <-client.Message:
			if err := c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				logger.Error("failed sending message", slog.Any("error", err))
				logger.Info("closing connection...")
				h.chatService.Logout(userID)
				return
			}
		}
	}
}
