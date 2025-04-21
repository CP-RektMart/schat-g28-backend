package chat

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

// @Summary			list chat
// @Tags			chats
// @Router			/api/v1/chats [GET]
// @Security		ApiKeyAuth
// @Success			200 {object}	dto.HttpResponse[ChatListResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleListChat(c *fiber.Ctx) error {
	ctx := c.UserContext()
	userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed get userID from context")
	}

	var directMessages []model.DirectMessage

	if err := h.db.
		Order("created_at").
		Joins("Receiver").
		Joins("Sender").
		Where("sender_id = ? OR receiver_id = ?", userID, userID).
		Find(&directMessages).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return apperror.NotFound("Messages not found", err)
		}
		return errors.Wrap(err, "failed getting sended messages")
	}

	var groupMessages []model.GroupMessage

	if err = h.db.
		Order("created_at").
		Joins("Group").
		Preload("Group").
		Find(&groupMessages).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return apperror.NotFound("Messages not found", err)
		}
		return errors.Wrap(err, "failed getting sended messages")
	}

	directs := h.toDirectMessageResponse(userID, directMessages)
	groups := h.toGroupMessageResponse(groupMessages)

	// return c.Status(200).JSON(fiber.Map{
	// 	"direct": directs,
	// 	"group":  groups,
	// })

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.ChatListResponse]{
		Result: dto.ChatListResponse{
			Directs: directs,
			Groups:  groups,
		},
	})
}

func (h *Handler) toDirectMessageResponse(userID uint, messages []model.DirectMessage) []dto.DirectChatListResponse {
	chats := make(map[uint]*dto.DirectChatListResponse)

	for _, msg := range messages {
		var talker model.User
		if msg.ReceiverID == userID {
			talker = msg.Sender
		} else {
			talker = msg.Receiver
		}

		chat, ok := chats[talker.ID]
		if !ok {
			chat = &dto.DirectChatListResponse{
				User:     dto.ToUserResponse(talker),
				Messages: []dto.DirectMessageResponse{},
			}
			chats[talker.ID] = chat
		}

		chat.Messages = append(chat.Messages, dto.ToDirectMessageResponse(msg))
	}

	return lo.MapToSlice(chats, func(_ uint, chat *dto.DirectChatListResponse) dto.DirectChatListResponse {
		return *chat
	})
}

func (h *Handler) toGroupMessageResponse(messages []model.GroupMessage) []dto.GroupChatListResponse {
	chats := make(map[uint]*dto.GroupChatListResponse)

	for _, msg := range messages {
		group := msg.Group

		chat, ok := chats[group.ID]
		if !ok {
			chat = &dto.GroupChatListResponse{
				Group:    dto.ToGroupListResponse(group),
				Messages: []dto.GroupMessageResponse{},
			}
			chats[group.ID] = chat
		}

		chat.Messages = append(chat.Messages, dto.ToGroupMessageResponse(msg))
	}

	return lo.MapToSlice(chats, func(_ uint, chat *dto.GroupChatListResponse) dto.GroupChatListResponse {
		return *chat
	})
}
