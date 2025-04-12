package user

import (
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"gorm.io/gorm"

	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

// @Summary			Update me
// @Description		Update user's profile
// @Tags			user
// @Router			/api/v1/me [PATCH]
// @Security		ApiKeyAuth
// @Param 			RequestBody 	body 	dto.UserUpdateRequest 	true 	"request request"
// @Success			200	{object}	dto.HttpResponse[dto.UserResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleUpdateMe(c *fiber.Ctx) error {
	ctx := c.UserContext()
	userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.UserUpdateRequest)
	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	updatedUser, err := h.updateUserDB(userID, req)
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}


	response := dto.UserResponse{
		ID:                updatedUser.ID,
		Name:              updatedUser.Name,
		Email:             updatedUser.Email,
		ProfilePictureURL: updatedUser.ProfilePictureURL,
	}

	return c.JSON(dto.HttpResponse[dto.UserResponse]{
		Result: response,
	})
}

func (h *Handler) updateUserDB(userID uint, req *dto.UserUpdateRequest) (*model.User, error) {
	var user model.User

	err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, "id = ?", userID).Error; err != nil {
			return errors.Wrap(err, "User not found")
		}

		updateField := func(field *string, newValue string) {
			if newValue != "" {
				*field = newValue
			}
		}

		updateField(&user.Name, req.Name)

		if err := tx.Save(&user).Error; err != nil {
			return errors.Wrap(err, "Failed to update user")
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to update user")
	}

	return &user, nil
}