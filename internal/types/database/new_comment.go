package database



type NewReqComment interface {
	GetPostID() int
	GetParentCommentID() int
	GetCommentContent() string
	GetCommentImage() string
	GetAuthorSession() (idSessionUser string)
}


