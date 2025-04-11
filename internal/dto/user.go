package dto

import "github.com/CP-RektMart/schat-g28-backend/internal/model"

type UserResponse struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	ProfilePictureURL string `json:"profilePictureUrl"`
}

type PublicUserResponse struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	ProfilePictureURL string `json:"profilePictureUrl"`
}

func ToUserResponse(user model.User) UserResponse {
	return UserResponse{
		ID:                user.ID,
		Name:              user.Name,
		Email:             user.Email,
		ProfilePictureURL: user.ProfilePictureURL,
	}
}

func ToPublicUserResponse(user model.User) PublicUserResponse {
	return PublicUserResponse{
		ID:                user.ID,
		Name:              user.Name,
		Email:             user.Email,
		ProfilePictureURL: user.ProfilePictureURL,
	}
}
