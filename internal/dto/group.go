package dto

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/samber/lo"
)

type CreateGroupRequest struct {
	Name              string  `json:"name"`
	ProfilePictureURL *string `json:"profilePictureUrl"`
	MemberIDs         []uint  `json:"memberIds"`
}

type UpdateGroupRequest struct {
	ID             uint    `params:"id"`
	Name           *string `json:"name"`
	ProfilePicture *string `json:"profilePicture"`
}

type DeleteGroupRequest struct {
	ID uint `params:"id"`
}

type JoinGroupRequest struct {
	ID uint `params:"id"`
}

type LeaveGroupRequest struct {
	ID uint `params:"id"`
}
type GetGroupReqest struct {
	ID uint `params:"id"`
}

type GroupListResponse struct {
	ID                uint         `json:"id"`
	ProfilePictureURL *string      `json:"profilePictureURL"`
	Name              string       `json:"name"`
	Owner             UserResponse `json:"owner"`
}

func ToGroupListResponse(g model.Group) GroupListResponse {
	return GroupListResponse{
		ID:                g.ID,
		ProfilePictureURL: g.ProfilePictureURL,
		Name:              g.Name,
		Owner:             ToUserResponse(g.Owner),
	}
}

func ToGroupListsResponse(groups []model.Group) []GroupListResponse {
	return lo.Map(groups, func(g model.Group, _ int) GroupListResponse {
		return ToGroupListResponse(g)
	})
}

type GroupDetailResponse struct {
	ID                uint                   `json:"id"`
	ProfilePictureURL *string                `json:"profilePictureURL"`
	Name              string                 `json:"name"`
	Owner             UserResponse           `json:"owner"`
	Members           []UserResponse         `json:"members"`
	Messages          []GroupMessageResponse `json:"messages"`
}

func ToGroupDetailReponse(g model.Group) GroupDetailResponse {
	return GroupDetailResponse{
		ID:                g.ID,
		ProfilePictureURL: g.ProfilePictureURL,
		Name:              g.Name,
		Owner:             ToUserResponse(g.Owner),
		Members:           ToUsersReponse(g.Members),
		Messages:          ToGroupMessagesResponse(g.Messages),
	}
}

type AddGroupMemberRequest struct {
	GroupID uint `params:"groupID"`
	UserID  uint `params:"userID"`
}

type KickGroupMemberRequest struct {
	GroupID uint `params:"groupID"`
	UserID  uint `params:"userID"`
}
