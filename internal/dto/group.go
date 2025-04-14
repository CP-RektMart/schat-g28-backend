package dto

type CreateGroupRequest struct {
	Name              string  `json:"name"`
	ProfilePictureURL *string `json:"profilePictureUrl"`
}

type UpdateGroupRequest struct {
	Name           string  `json:"name"`
	ProfilePicture *string `json:"profilePicture"`
}
