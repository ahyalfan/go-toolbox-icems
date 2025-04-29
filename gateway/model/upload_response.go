package model

type UploadResponse struct {
	ID      string `json:"id"`
	UrlPath string `json:"url_path"`
}

func (u *UploadResponse) GetID() string {
	return u.ID
}
