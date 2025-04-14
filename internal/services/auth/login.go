package auth

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			Login
// @Tags			auth
// @Router			/api/v1/auth/login [POST]
// @Param 			RequestBody 	body 	dto.LoginRequest 	true 	"request request"
// @Success			200	{object}	dto.HttpResponse[dto.LoginResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleLogin(c *fiber.Ctx) error {
	ctx := c.Context()

	req := new(dto.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}

	OAuthUser, err := h.googleOauth.ValidateIDToken(ctx, req.IDToken)
	if err != nil {
		return errors.Wrap(err, "failed to validate id token")
	}

	user := h.getUserByEmail(OAuthUser.Email)
	if user == nil {
		if user, err = h.createUser(OAuthUser); err != nil {
			return errors.Wrap(err, "failed create user")
		}
	}

	token, err := h.jwtService.GenerateAndStoreTokenPair(ctx, user)
	if err != nil {
		return errors.Wrap(err, "failed to generate token pair")
	}

	result := dto.LoginResponse{
		TokenResponse: dto.ToTokenResponse(*token),
		User:          dto.ToUserResponse(*user),
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.LoginResponse]{
		Result: result,
	})
}

func (h *Handler) createUser(user model.User) (*model.User, error) {
	if err := h.store.DB.Save(&user).Error; err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}
	return &user, nil
}

func (h *Handler) getUserByEmail(email string) *model.User {
	var user model.User
	if err := h.store.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil
	}
	return &user
}
