package chat

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
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

	var groupMessage []model.GroupMessage

	if err = h.db.
		Order("created_at").
		// Joins("Group").
		// Joins("Sender").
		Where("sender_id = ?", userID).
		Find(&groupMessage).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return apperror.NotFound("Messages not found", err)
		}
		return errors.Wrap(err, "failed getting sended messages")
	}

	// return c.Status(fiber.StatusOK).JSON(dto.HttpListResponse[dto.ChatResponse]{
	// 	Result: messages,
	// })
	return c.Status(200).JSON(fiber.Map{
		"direct_messages": directMessages,
		"group_messages":  groupMessage,
	})
}
