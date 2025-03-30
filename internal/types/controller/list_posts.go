package controller

type ListPostsResp interface {
	GetList() []ItemPostsResp
}

type ItemPostsResp interface {
	GetPostID() string
	GetTitle() string
	GetPostContent() string
	GetPostImageURL() string
	GetPostTime() string
}
