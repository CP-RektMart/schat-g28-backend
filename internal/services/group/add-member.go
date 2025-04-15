package group

import "github.com/gofiber/fiber/v2"

// @Summary			add group member
// @Tags			groups
// @Router			/api/v1/groups/{groupID}/members/{userID} [POST]
// @Security		ApiKeyAuth
// @Success			200
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleAddGroupMember(c *fiber.Ctx) error {
	return nil
}
