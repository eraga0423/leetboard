package controller

import "io"

type NewPostReq interface {
	GetTitle() string
	GetPostContent() string
	GetImage() io.Reader
	GetName() string
	GetAuthorIDSession() string
}

type NewPostResp interface {
}
