package dtos

type CreateShortLinkRequest struct {
	OriginalUrl string `json:"original_url" validate:"required,url"`
}
