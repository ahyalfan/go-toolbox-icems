package model

import "mime/multipart"

type Request interface {
	GetID() string
}

type RequestMultiForm interface {
	GetID() string
	GetFileName() string
	GetFile() *multipart.FileHeader
}
