package dto

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/samber/lo"
)

type UserUpdateRequest struct {
	Name              *string `json:"name"`
	Email             *string `json:"email"`
	ProfilePictureURL *string `json:"profilePictureUrl"`
	Color             *string `json:"color"`
}

type UserResponse struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	ProfilePictureURL string `json:"profilePictureUrl"`
}

type GetUserReqest struct {
	ID uint `params:"id"`
}

func ToUserResponse(user model.User) UserResponse {
	return UserResponse{
		ID:                user.ID,
		Name:              user.Name,
		Email:             user.Email,
		ProfilePictureURL: user.ProfilePictureURL,
	}
}

func ToUsersReponse(users []model.User) []UserResponse {
	return lo.Map(users, func(u model.User, _ int) UserResponse {
		return ToUserResponse(u)
	})
}
