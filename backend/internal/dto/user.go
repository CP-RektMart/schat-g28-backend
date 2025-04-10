package dto

import "github.com/CP-RektMart/computer-network-g28/backend/internal/model"

type UserResponse struct {
	ID                uint           `json:"id"`
	Name              string         `json:"name"`
	Email             string         `json:"email"`
	PhoneNumber       string         `json:"phoneNumber"`
	ProfilePictureURL string         `json:"profilePictureUrl"`
	Role              model.UserRole `json:"role"`
	Facebook          string         `json:"facebook,omitempty"`
	Instagram         string         `json:"instagram,omitempty"`
	Bank              string         `json:"bank,omitempty"`
	AccountNo         string         `json:"accountNo,omitempty"`
	BankBranch        string         `json:"bankBranch,omitempty"`
}

func ToUserResponse(user model.User) UserResponse {
	return UserResponse{
		ID:                user.ID,
		Name:              user.Name,
		Email:             user.Email,
		ProfilePictureURL: user.ProfilePictureURL,
	}
}
