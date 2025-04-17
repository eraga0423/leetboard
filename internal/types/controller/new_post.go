package controller

import (
	"mime/multipart"
)

type NewPostReq interface {
	GetTitle() string
	GetPostContent() string
	GetImage() ItemMetaData
	GetName() string
	GetAuthorIDSession() string
}

type ItemMetaData interface {
	GetFileIO() multipart.File
	GetObjectSize() int64
	GetContentType() string
}

type NewPostResp interface{}
