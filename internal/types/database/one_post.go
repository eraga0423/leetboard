package database

import "time"

type OnePostResp interface {
	// post
	GetComments() []Comment
}

type Comment interface {
	GetParent() OneComment
	GetChildren() []OneComment
}
type OneComment interface {
	GetCommentID() int
	GetPostID() int
	GetCommentContent() string
	GetCommentImage() string
	GetCommentTime() time.Time
}
type OnePostReq interface {
	ReqPostID() int
}
