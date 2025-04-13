package database

import "time"

// for archive - post
type ListPostsArchiveResp interface {
	GetArchiveList() []ItemPostsArchiveResp
}

type ItemPostsArchiveResp interface {
	GetPostID() int
	GetTitle() string
	GetPostContent() string
	GetPostImageURL() string
	GetPostTime() time.Time
}
