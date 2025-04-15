package friend

import "github.com/gofiber/fiber/v2"

// @Summary			list user's friends
// @Tags			friends
// @Router			/api/v1/friends [GET]
// @Security		ApiKeyAuth
// @Success			200 {object}	dto.HttpListResponse[dto.UserResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleListFriends(c *fiber.Ctx) error {
	return nil
}
