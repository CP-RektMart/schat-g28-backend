package dto

import "github.com/CP-RektMart/schat-g28-backend/internal/model"

type FileDetailResponse struct {
	ID      uint    `json:"id"`
	URL     string  `json:"url"`
	OwnerID uint    `json:"ownerId"`
	Path    *string `json:"path"`
}

type DeleteFileRequest struct {
	ID uint `params:"id"`
}

func ToFileDetailResponse(f model.File) FileDetailResponse {
	return FileDetailResponse{
		ID:      f.ID,
		URL:     f.URL,
		OwnerID: f.OwnerID,
		Path:    f.Path,
	}
}
