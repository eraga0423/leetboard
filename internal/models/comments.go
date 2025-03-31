package models

import "time"

type MonoComment struct {
	CommentID      int
	PostID         int
	CommentContent string
	CommentImage   string
	CommentTime    time.Time
}

type SubComments []MonoComment

type Comment struct {
	ParentComment MonoComment
	SubComments
}
type Comments []Comment
