package model

import "mime/multipart"

type UploadRequest struct {
	ID         string
	BucketName string `form:"bucket_name"`
	PathFile   string `form:"path_file"`
	File       *multipart.FileHeader
}

func (u *UploadRequest) GetID() string {
	return u.ID
}

func (u *UploadRequest) GetFileName() string {
	return u.File.Filename
}

func (u *UploadRequest) GetFile() *multipart.FileHeader {
	return u.File
}
