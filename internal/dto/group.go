package dto

type CreateGroupRequest struct {
	Name           string `json:"name"`
	ProfilePicture string `json:"profilePicture"`
}

type UpdateGroupRequest struct {
	Name           string `json:"name"`
	ProfilePicture string `json:"profilePicture"`
}
