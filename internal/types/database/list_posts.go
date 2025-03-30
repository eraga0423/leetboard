package database

import "time"

type ListPostsResp interface {
	GetList() []ItemPostsResp
}

type ItemPostsResp interface {
	GetPostID() string
	GetTitle() string
	GetPostContent() string
	GetPostImageURL() string
	GetPostTime() time.Time
	GetComment() []Comment
}

type Comment interface {
	GetContent() string
	GetComment() string
	GetCommentImageURL() string
	GetCommentTime() time.Time
}
