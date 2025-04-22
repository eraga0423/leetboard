package controller

import "time"

type ListArchivePostsResp interface {
	GetList() []ItemArchivePostsResp
}

type ItemArchivePostsResp interface {
	GetPostID() int
	GetTitle() string
	GetPostContent() string
	GetPostImageURL() string
	GetPostTime() time.Time
}
