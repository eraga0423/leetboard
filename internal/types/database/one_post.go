package database

import "time"

type OnePostResp interface {
	// post
	GetComments(idPost int) []Comment
}

type Comment interface {
	GetParentComment(idComment int) (OneComment, error)
	GetSubComment(idParentComment int) []OneComment
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
