package dto

type DirectChatListResponse struct {
	User        UserResponse          `json:"user"`
	LastMessage DirectMessageResponse `json:"lastMessage"`
}

type GroupChatListResponse struct {
	Group       GroupListResponse      `json:"group"`
	LastMessage []GroupMessageResponse `json:"lastMessage"`
}

type ChatListResponse struct {
	Directs []DirectChatListResponse `json:"directs"`
	Groups  []GroupChatListResponse  `json:"groups"`
}

type DirectChatDetailResponse struct {
	User     UserResponse            `json:"user"`
	Messages []DirectMessageResponse `json:"messages"`
}
