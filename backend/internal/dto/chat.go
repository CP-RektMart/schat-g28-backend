package dto

type ChatDirectResponse struct {
	User     PublicUserResponse        `json:"user"`
	Messages []DirectMessageResponse `json:"messages"`
}

type ChatGroupResponse struct {
	User     PublicUserResponse        `json:"user"`
	Messages []GroupMessageResponse `json:"messages"`
}
