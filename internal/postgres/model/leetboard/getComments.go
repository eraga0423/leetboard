package leetboard

import (
	"1337b0rd/internal/types/database"
	"time"
)

type onecomment struct {
	commentID      int
	postID         int
	commentContent string
	commentImage   string
	commentTime    time.Time
}
type comment struct {
	parentComment onecomment
	subComments   []onecomment
}
type allcomment struct {
	comments []comment
}

func (a allcomment) GetComments() []database.Comment {
	com := make([]database.Comment, len(a.comments))
	for i, _ := range a.comments {
		com[i] = a.comments[i]
	}
	return com

}
func (c comment) GetSubComment() []database.OneComment {
	sub := make([]database.OneComment, len(c.subComments))
	for i, s := range c.subComments {
		sub[i] = s
	}
	return sub
}

func (c comment) GetParentComment() database.OneComment {
	return c.parentComment
}
func (m onecomment) GetCommentID() int {
	return m.commentID
}
func (m onecomment) GetPostID() int {
	return m.postID
}
func (m onecomment) GetCommentContent() string {
	return m.commentContent
}
func (m onecomment) GetCommentImage() string {
	return m.commentImage
}
func (m onecomment) GetCommentTime() time.Time {
	return m.commentTime
}
