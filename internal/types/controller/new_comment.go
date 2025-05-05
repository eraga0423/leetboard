package controller

import "mime/multipart"

type NewCommentReq interface {
	GetAvatarImageURL() string
	GetAvatarName() string
	GetParentCommentID() string
	GetSessionID() string
	GetContent() string
	GetImageComment() MetaDataComment
	GetPostID() string
}
type MetaDataComment interface {
	GetFileIO() multipart.File
	GetObjectSize() int64
	GetContentType() string
}
type NewCommentResp interface{}
