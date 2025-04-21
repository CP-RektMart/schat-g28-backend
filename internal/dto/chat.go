package dto

type DirectChatListResponse struct {
	User     UserResponse            `json:"user"`
	Messages []DirectMessageResponse `json:"messages"`
}

type GroupChatListResponse struct {
	Group    GroupListResponse      `json:"group"`
	Messages []GroupMessageResponse `json:"messages"`
}

type ChatListResponse struct {
	Directs []DirectChatListResponse `json:"directs"`
	Groups  []GroupChatListResponse  `json:"groups"`
}

type DirectChatDetailResponse struct {
	User     UserResponse            `json:"user"`
	Messages []DirectMessageResponse `json:"messages"`
}
