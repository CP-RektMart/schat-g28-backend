package group

import "github.com/gofiber/fiber/v2"

// @Summary			create group
// @Tags			groups
// @Router			/api/v1/groups [POST]
// @Security		ApiKeyAuth
// @Param 			groupDetail 	body 	dto.CreateGroupRequest 	true  "group detail"
// @Success			200
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleCreateGroup(c *fiber.Ctx) error {
	return nil
}
