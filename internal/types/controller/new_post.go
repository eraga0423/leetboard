package controller

type NewPostReq interface {
	GetTitle() string
	GetPostContent() string
	GetImage() []byte
	GetName() string
	GetAuthorID() string
}

type NewPostResp interface {
}
