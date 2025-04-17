package controller

import "io"

type NewPostReq interface {
	GetTitle() string
	GetPostContent() string
	GetImage() ItemMetaData
	GetName() string
	GetAuthorIDSession() string
}

type ItemMetaData interface {
	GetFileIO() io.Reader
	GetObjectSize() int64
	GetContentType() string
}

type NewPostResp interface {
}
