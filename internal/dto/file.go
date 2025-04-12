package dto

type FileDetailResponse struct {
	ID      uint   `json:"id"`
	URL     string `json:"url"`
	OwnerID uint   `json:"ownerId"`
}