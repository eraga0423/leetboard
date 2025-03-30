package database

type NewPostReq interface {
	GetTitle() string
	GetPostContent() string
	GetImage() string
}

type NewPostResp interface{}
