package chat

import "github.com/gofiber/fiber/v2"

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
	return nil
}