package message

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

// @Summary      list all direct message chats
// @Description  list all direct message chats for each individual user
// @Tags         direct_message
// @Router       /api/v1/messages [GET]
// @Security	 ApiKeyAuth
// @Success      200    {object}  dto.HttpListResponse[dto.ChatDirectResponse]
// @Failure      401    {object}  dto.HttpError
// @Failure      404    {object}  dto.HttpError
// @Failure      500    {object}  dto.HttpError
func (h *Handler) HandleListMessages(c *fiber.Ctx) error {
	userID, err := h.authentication.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed getting userID from context")
	}

	var messages []model.DirectMessage

	if err := h.store.DB.
		Order("created_at").
		Joins("Receiver").
		Joins("Sender").
		Where("sender_id = ? OR receiver_id = ?", userID, userID).
		Find(&messages).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return apperror.NotFound("Messages not found", err)
		}
		return errors.Wrap(err, "failed getting sended messages")
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpListResponse[dto.ChatDirectResponse]{
		Result: h.toDirectChatResponse(userID, messages),
	})
}

func (h *Handler) toDirectChatResponse(userID uint, messages []model.DirectMessage) []dto.ChatDirectResponse {
	chats := make(map[uint]*dto.ChatDirectResponse)
	var talker model.User

	for _, msg := range messages {
		if msg.ReceiverID == userID {
			talker = msg.Sender
		} else {
			talker = msg.Receiver
		}

		chat, ok := chats[talker.ID]
		if !ok {
			chats[talker.ID] = &dto.ChatDirectResponse{
				// User:     dto.ToPublicUserResponse(talker),
				Messages: make([]dto.DirectMessageResponse, 0),
			}
			chat = chats[talker.ID]
		}

		chat.Messages = append(chat.Messages, dto.ToDirectMessageResponse(msg))
	}

	return lo.MapToSlice(chats, func(_ uint, chat *dto.ChatDirectResponse) dto.ChatDirectResponse {
		return *chat
	})
}
