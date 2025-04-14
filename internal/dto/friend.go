package dto

type AddFriendRequest struct {
	FriendID uint `params:"friendID"`
}

type UnFriendRequest struct {
	FriendID uint `params:"friendID"`
}
