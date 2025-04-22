package controller

import "time"

type ArchiveOnePostReq interface {
	GetPostID() int
}

type ArchiveOnePostResp interface {
	GetOnePost() ArchiveRespOnePost
	GetComments() []ArchiveComment
}

type ArchiveComment interface {
	GetParent() ArchiveOneComment
	GetChildren() []ArchiveOneComment
}
type ArchiveOneComment interface {
	GetCommentID() int
	GetPostID() int
	GetAuthor() ArchiveRespCommentAuthor
	GetCommentContent() string
	GetCommentImage() string
	GetCommentTime() time.Time
}
type ArchiveRespOnePost interface {
	GetTitle() string
	GetPostContent() string
	GetPostUrlImage() string
	GetPostTime() time.Time
	GetAuthorPost() ArchiveRespOnePostAuthor
}

type ArchiveRespOnePostAuthor interface {
	GetName() string
	GetImageURL() string
	GetSessionID() string
}
type ArchiveRespCommentAuthor interface {
	GetName() string
	GetImageURL() string
	GetSessionID() string
}
