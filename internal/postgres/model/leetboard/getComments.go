package leetboard

import (
	"time"

	"1337b0rd/internal/types/database"
)

type OneCommentData struct {
	ID      int
	PostID  int
	Content string
	Image   string
	Time    time.Time
}
type CommentNode struct {
	Parent   OneCommentData
	Children []OneCommentData
}

type OnePostResponse struct {
	Comments []CommentNode
}

func (l *Leetboard) OnePost(r database.OnePostReq) (database.OnePostResp, error) {
	idPost := r.ReqPostID()
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

	commentMap := map[int]*CommentNode{}
	for rows.Next() {
		var parentID, parentPostID int
		var parentContent, parentImage string
		var parentTime time.Time

		var childID, childPostID int
		var childContent, childImage string
		var childTime time.Time

		err := rows.Scan(
			&parentID, &parentPostID, &parentContent, &parentImage, &parentTime,
			&childID, &childPostID, &childContent, &childImage, &childTime,
		)
		if err != nil {
			return nil, err
		}

		if _, ok := commentMap[parentID]; !ok {
			commentMap[parentID] = &CommentNode{
				Parent: OneCommentData{
					ID:      parentID,
					PostID:  parentPostID,
					Content: parentContent,
					Image:   parentImage,
					Time:    parentTime,
				},
			}
		}

		if childID != 0 {
			commentMap[parentID].Children = append(commentMap[parentID].Children, OneCommentData{
				ID:      childID,
				PostID:  childPostID,
				Content: childContent,
				Image:   childImage,
				Time:    childTime,
			})
		}
	}

	var commentList []CommentNode
	for _, v := range commentMap {
		commentList = append(commentList, *v)
	}

	return OnePostResponse{Comments: commentList}, nil
}

func (r OnePostResponse) GetComments() []database.Comment {
	comments := make([]database.Comment, len(r.Comments))
	for num, comment := range r.Comments {
		comments[num] = comment
	}
	return comments
}
func (c CommentNode) GetParent() database.OneComment {
	return c.Parent
}

func (c CommentNode) GetChildren() []database.OneComment {
	result := make([]database.OneComment, len(c.Children))
	for i := range c.Children {
		result[i] = c.Children[i]
	}
	return result
}

func (c OneCommentData) GetCommentID() int         { return c.ID }
func (c OneCommentData) GetPostID() int            { return c.PostID }
func (c OneCommentData) GetCommentContent() string { return c.Content }
func (c OneCommentData) GetCommentImage() string   { return c.Image }
func (c OneCommentData) GetCommentTime() time.Time { return c.Time }
