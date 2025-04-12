package dto

type ChatDirectResponse struct {
	User     UserResponse            `json:"user"`
	Messages []DirectMessageResponse `json:"messages"`
}

type ChatGroupResponse struct {
	User     UserResponse           `json:"user"`
	Messages []GroupMessageResponse `json:"messages"`
}
