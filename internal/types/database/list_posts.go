package database

import "time"

// for catalog - post
type ListPostsResp interface {
	GetList() []ItemPostsResp
}

type ItemPostsResp interface {
	GetPostID() int
	GetTitle() string
	GetPostContent() string
	GetPostImageURL() string
	GetPostTime() time.Time
}
