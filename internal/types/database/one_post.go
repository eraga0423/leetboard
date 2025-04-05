package database

import "time"

type OnePost interface {
	//post
	GetComments() []Comment
}

type Comment interface {
	GetParentComment() OneComment
	GetSubComment() []OneComment
}
type OneComment interface {
	GetCommentID() int
	GetPostID() int
	GetCommentContent() string
	GetCommentImage() string
	GetCommentTime() time.Time
}
