package group

import "github.com/gofiber/fiber/v2"

// @Summary			update group
// @Tags			groups
// @Router			/api/v1/groups/{groupID} [PATCH]
// @Security		ApiKeyAuth
// @Param 			groupID 	path 	uint 	true  "group id"
// @Param 			groupDetail 	body 	dto.UpdateGroupRequest 	true  "group detail"
// @Success			200
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleUpdateGroup(c *fiber.Ctx) error {
	return nil
}
