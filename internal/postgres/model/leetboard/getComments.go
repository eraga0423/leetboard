package leetboard

import (
	"time"

	"1337b0rd/internal/types/database"
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

func (l Leetboard) GetComments(idPost int) ([]database.Comment, error) {
	rows, err := l.db.Query(`
	SELECT 
    c.comment_id,
    c.post_id,
    c.comment_content,
    c.comment_image,
    c.comment_time,
    s.comment_child,
	sub.post_id AS child_comment_post,
    sub.comment_content AS child_content,
    sub.comment_image AS child_image,
    sub.comment_time AS child_time
FROM comments c
LEFT JOIN subcomments s ON s.comment_parent = c.comment_id
LEFT JOIN comments sub ON sub.comment_id = s.comment_child
WHERE c.post_id = $1
`, idPost)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []database.Comment
	commentsMap := make(map[onecomment][]onecomment)
	for rows.Next() {
		var comment onecomment
		var childComID, childPostID int
		var childComContent, childComImage string
		var childComTime time.Time
		err := rows.Scan(
			&comment.commentID,
			&comment.postID,
			&comment.commentContent,
			&comment.commentImage,
			&comment.commentTime,
			&childComID,
			&childPostID,
			&childComContent,
			&childComImage,
			&childComTime,
		)
		if err != nil {
			return nil, err
		}
		value, exist := commentsMap[comment]
		if exist {
			value = append(value, onecomment{
				commentID:      childComID,
				postID:         childPostID,
				commentContent: childComContent,
				commentImage:   childComImage,
				commentTime:    childComTime,
			},
			)
		} else {
			value = []onecomment{
				{
					commentID:      childComID,
					postID:         childPostID,
					commentContent: childComContent,
					commentImage:   childComImage,
					commentTime:    childComTime,
				},
			}
		}
		commentsMap[comment] = value

	}

	com := make([]database.Comment, len(a.comments))
	for i := range a.comments {
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
