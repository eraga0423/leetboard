package database

import "time"

type OnePostResp interface {
	GetOnePost() RespOnePost
	GetComments() []Comment
}

type Comment interface {
	GetParent() OneComment
	GetChildren() []OneComment
}
type OneComment interface {
	GetCommentID() int
	GetPostID() int
	GetAuthor() RespCommentAuthor
	GetCommentContent() string
	GetCommentImage() string
	GetCommentTime() time.Time
}
type RespOnePost interface {
	GetTitle() string
	GetPostContent() string
	GetPostUrlImage() string
	GetPostTime() time.Time
	GetAuthorPost() RespOnePostAuthor
}

type RespOnePostAuthor interface {
	GetName() string
	GetImageURL() string
	GetSessionID() string
}
type RespCommentAuthor interface {
	GetName() string
	GetImageURL() string
	GetSessionID() string
}

type OnePostReq interface {
	ReqPostID() int
}
