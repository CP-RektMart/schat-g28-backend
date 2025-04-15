package chat

import "github.com/gofiber/fiber/v2"

// @Summary			get dm chats detail
// @Tags			chats
// @Router			/api/v1/chats/friends/{friendID} [GET]
// @Security		ApiKeyAuth
// @Success			200 {object}	dto.HttpResponse[DirectChatDetailResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleGetDMChat(c *fiber.Ctx) error {
	return nil
}
