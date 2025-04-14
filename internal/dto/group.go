package dto

import "github.com/CP-RektMart/schat-g28-backend/internal/model"

type CreateGroupRequest struct {
	Name              string  `json:"name"`
	ProfilePictureURL *string `json:"profilePictureUrl"`
}

type UpdateGroupRequest struct {
	Name           string  `json:"name"`
	ProfilePicture *string `json:"profilePicture"`
}

type GetGroupReqest struct {
	ID uint `params:"id"`
}

type GroupResponse struct {
	ID                uint                   `json:"id"`
	ProfilePictureURL *string                `json:"profilePictureURL"`
	Name              string                 `json:"name"`
	Owner             UserResponse           `json:"owner"`
	Members           []UserResponse         `json:"members"`
	Messages          []GroupMessageResponse `json:"messages"`
}

func ToGroupReponse(g model.Group) GroupResponse {
	return GroupResponse{
		ID:                g.ID,
		ProfilePictureURL: g.ProfilePictureURL,
		Name:              g.Name,
		Owner:             ToUserResponse(g.Owner),
		Members:           ToUsersReponse(g.Members),
		Messages:          ToGroupMessagesResponse(g.Messages),
	}
}
