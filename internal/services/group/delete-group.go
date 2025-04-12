package group

import "github.com/gofiber/fiber/v2"

// @Summary			delete group
// @Tags			groups
// @Router			/api/v1/groups/{groupID} [DELETE]
// @Security		ApiKeyAuth
// @Param 			groupID 	path 	uint 	true  "group id"
// @Success			200
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleDeleteGroup(c *fiber.Ctx) error {
	return nil
}
