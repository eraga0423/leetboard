package controller

import "time"

type ListPostsResp interface {
	GetList() []ItemPostsResp
}

type ItemPostsResp interface {
	GetPostID() int
	GetTitle() string
	GetPostContent() string
	GetPostImageURL() string
	GetPostTime() time.Time
	GetComments() []ItemComment
	GetUser() User
}

type ItemComment interface {
	GetParentComment() []ItemMonoComment
	GetSubComment()
}

type ItemSubComment interface {
	GetSubComment() []ItemMonoComment
}

type ItemMonoComment interface {
	GetCommentID() int
	GetPostID() int
	GetCommentContent() string
	GetCommentImage() string
	GetCommentTime() time.Time
}

type User interface {
	GetUserID() int
	GetUserName() string
	GetAvatar() string
}
