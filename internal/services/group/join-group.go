package group

import "github.com/gofiber/fiber/v2"

// @Summary			join group
// @Tags			groups
// @Router			/api/v1/groups/{groupID}/join [GET]
// @Security		ApiKeyAuth
// @Param 			groupID 	path 	uint 	true  "group id"
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleJoinGroup(c *fiber.Ctx) error {
	return nil
}
