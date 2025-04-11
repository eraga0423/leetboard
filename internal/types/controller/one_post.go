package controller

import "time"

type OnePostReq interface {
	GetPostID() int
}
type OnePostResp interface {
	GetOnePost() ItemOnePost
	GetComments() []Comment
}

type Comment interface {
	GetParent() OneComment
	GetChildren() []OneComment
}
type OneComment interface {
	GetCommentID() int
	GetPostID() int
	GetAuthor() Author
	GetCommentContent() string
	GetCommentImage() string
	GetCommentTime() time.Time
}

type Author interface {
	GetName() string
	GetImageURL() string
	GetSessionID() string
}

type ItemOnePost interface {
	GetTitle() string
	GetContent() string
	GetImageURL() string
	GetPostTime() time.Time
	GetAuthorPost() Author
}
