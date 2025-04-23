package chat

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"sync"

	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/pkg/logger"
	"github.com/cockroachdb/errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type EventType string

const (
	EventError   EventType = "ERROR"
	EventMessage EventType = "MESSAGE"
	EventGroup   EventType = "GROUP"
)

type Client struct {
	Message   chan string
	Terminate chan bool
}

type Server struct {
	clients  map[uint]*Client
	db       *gorm.DB
	validate *validator.Validate
}

var (
	instance *Server
	once     sync.Once
)

func NewServer(db *gorm.DB, validate *validator.Validate) *Server {
	once.Do(func() {
		instance = &Server{
			db:       db,
			validate: validate,
			clients:  make(map[uint]*Client),
		}
	})

	return instance
}

func (c *Server) Register(userID uint) *Client {
	client := Client{
		Message:   make(chan string),
		Terminate: make(chan bool),
	}
	c.clients[userID] = &client
	return &client
}

func (c *Server) Logout(userID uint) {
	if c.isUserExist(userID) {
		c.clients[userID].Terminate <- true
		delete(c.clients, userID)
	}
}

func (c *Server) SendDirectRawString(senderID uint, msg string) {
	var msgReq dto.DirectMessageRequest
	if err := json.Unmarshal([]byte(msg), &msgReq); err != nil {
		logger.Error("Failed Unmarshal json", slog.Any("error", err))
		c.sendMessage(EventError, senderID, "invalid message")
		return
	}

	if err := c.validate.Struct(msgReq); err != nil {
		logger.Error("Failed validate message request", slog.Any("error", err))
		c.sendMessage(EventError, senderID, "invalid message")
		return

	}

	msgModel := dto.ToDirectMessageModel(senderID, msgReq)
	if msgModel.ReceiverID == senderID {
		logger.Error("cannot send message to yourself")
		c.sendMessage(EventError, senderID, "cannot send message to yourself")
		return
	}

	if err := c.db.Create(&msgModel).Error; err != nil {
		logger.Error("failed inserting message to store", slog.Any("error", err))
		c.sendMessage(EventError, senderID, "internal error")
		return
	}

	json, err := json.Marshal(dto.ToDirectMessageResponse(msgModel))
	if err != nil {
		logger.Error("failed Marshal realtime message response to json", slog.Any("error", err))
		c.sendMessage(EventError, senderID, "internal error")
		return
	}

	c.sendMessage(EventMessage, msgModel.ReceiverID, string(json))
	c.sendMessage(EventMessage, msgModel.SenderID, string(json))
}

func (c *Server) SendGroupRawString(senderID uint, msg string) {
	var msgReq dto.GroupMessageRequest
	if err := json.Unmarshal([]byte(msg), &msgReq); err != nil {
		logger.Error("Failed Unmarshal json", slog.Any("error", err))
		c.sendMessage(EventError, senderID, "invalid message")
		return
	}

	if err := c.validate.Struct(msgReq); err != nil {
		logger.Error("Failed validate message request", slog.Any("error", err))
		c.sendMessage(EventError, senderID, "invalid message")
		return

	}

	msgModel := dto.ToGroupMessageModel(senderID, msgReq)

	if err := c.db.Create(&msgModel).Error; err != nil {
		logger.Error("failed inserting message to store", slog.Any("error", err))
		c.sendMessage(EventError, senderID, "internal error")
		return
	}

	json, err := json.Marshal(dto.ToGroupMessageResponse(msgModel))
	if err != nil {
		logger.Error("failed Marshal realtime message response to json", slog.Any("error", err))
		c.sendMessage(EventError, senderID, "internal error")
		return
	}

	// Broadcast the message to all group members
	c.broadcastToGroup(msgModel.GroupID, string(json), senderID)
}

func (c *Server) broadcastToGroup(groupID uint, msg string, senderID uint) {
	// Fetch group members from the database
	var group model.Group
	if err := c.db.Preload("Owner").Preload("Members").First(&group, groupID).Error; err != nil {
		logger.Error("Failed to fetch group members", slog.Any("error", err))
		c.sendMessage(EventError, senderID, "group not found")
		return
	}

	// group.Members = append(group.Members, group.Owner)
	fmt.Println(group.Members)

	// Send the message to all connected members of the group
	for _, member := range group.Members {
		if c.isUserExist(member.ID) {
			c.sendMessage(EventGroup, member.ID, msg)
		}
	}
}

func (c *Server) SendMessageModel(msg model.DirectMessage) error {
	if err := c.db.Create(&msg).Error; err != nil {
		return errors.Wrap(err, "failed save massage record to store")
	}

	json, err := json.Marshal(dto.ToDirectMessageResponse(msg))
	if err != nil {
		return errors.Wrap(err, "failed Marshal to json")
	}

	c.sendMessage(EventMessage, msg.ReceiverID, string(json))
	c.sendMessage(EventMessage, msg.SenderID, string(json))

	return nil
}

func (c *Server) sendMessage(event EventType, receiverID uint, msg string) {
	msg = fmt.Sprintf("%s %s", event, msg)
	if c.isUserExist(receiverID) {
		c.clients[receiverID].Message <- msg
	}
}

func (c *Server) isUserExist(userID uint) bool {
	_, ok := c.clients[userID]
	return ok
}
