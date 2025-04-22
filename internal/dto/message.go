package dto

import (
	"time"

	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/samber/lo"
)

// Requests
type DirectMessageRequest struct {
	Content    string `json:"content" validate:"required"`
	ReceiverID uint   `json:"receiverId" validate:"required"`
}

type GroupMessageRequest struct {
	Content string `json:"content" validate:"required"`
	GroupID uint   `json:"groupId" validate:"required"`
}

// Responses
type DirectMessageResponse struct {
	ID         uint      `json:"id"`
	Content    string    `json:"content"`
	ReceiverID uint      `json:"receiverId"`
	SenderID   uint      `json:"senderId"`
	SendedAt   time.Time `json:"sendedAt"`
}

type GroupMessageResponse struct {
	ID       uint      `json:"id"`
	Content  string    `json:"content"`
	GroupID  uint      `json:"groupId"`
	SenderID uint      `json:"senderId"`
	SendedAt time.Time `json:"sendedAt"`
}

// To function
func ToDirectMessageModel(senderID uint, message DirectMessageRequest) model.DirectMessage {
	return model.DirectMessage{
		Content:    message.Content,
		SenderID:   senderID,
		ReceiverID: message.ReceiverID,
	}
}

func ToDirectMessageResponse(message model.DirectMessage) DirectMessageResponse {
	return DirectMessageResponse{
		ID:         message.ID,
		Content:    message.Content,
		SenderID:   message.SenderID,
		ReceiverID: message.ReceiverID,
		SendedAt:   message.CreatedAt,
	}
}

func ToGroupMessageModel(senderID uint, message GroupMessageRequest) model.GroupMessage {
	return model.GroupMessage{
		Content:  message.Content,
		SenderID: senderID,
		GroupID:  message.GroupID,
	}
}

func ToGroupMessageResponse(message model.GroupMessage) GroupMessageResponse {
	return GroupMessageResponse{
		ID:       message.ID,
		Content:  message.Content,
		SenderID: message.SenderID,
		GroupID:  message.GroupID,
		SendedAt: message.CreatedAt,
	}
}

func ToDirectMessagesResponse(messages []model.DirectMessage) []DirectMessageResponse {
	return lo.Map(messages, func(msg model.DirectMessage, _ int) DirectMessageResponse {
		return ToDirectMessageResponse(msg)
	})
}

func ToGroupMessagesResponse(messages []model.GroupMessage) []GroupMessageResponse {
	return lo.Map(messages, func(msg model.GroupMessage, _ int) GroupMessageResponse {
		return ToGroupMessageResponse(msg)
	})
}
