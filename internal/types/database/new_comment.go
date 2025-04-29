package database



type NewReqComment interface {
	GetAuthorName() string
	GetAuthorAvatarURL()string
	GetPostID() int
	GetParentCommentID() int
	GetCommentContent() string
	GetCommentImage() string
	GetAuthorSession() (idSessionUser string)
}


