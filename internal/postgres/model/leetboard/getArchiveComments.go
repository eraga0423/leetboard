package leetboard

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"1337b0rd/internal/types/database"
)

type archiveOneCommentData struct {
	id      int
	postID  int
	content string
	image   string
	time    time.Time
	author  archiveOneCommentAuthor
}
type archiveOneCommentAuthor struct {
	authorName      string
	authorImageURL  string
	authorSessionID string
}

type archiveCommentNode struct {
	parent   archiveOneCommentData
	children []archiveOneCommentData
}

type archiveOnePostResponse struct {
	comments []archiveCommentNode
	post     archiveOnePost
}

func (l *Leetboard) OneArchivePost(ctx context.Context, r database.ArchiveOnePostReq) (database.ArchiveOnePostResp, error) {
	log := l.logger.With(slog.String("handler", "OneArchivePost"))

	idPost := r.ReqPostID()
	onePost, err := archiveReturnOnePost(idPost, l.db)
	if err != nil {
		log.ErrorContext(ctx, "Error getting one post", slog.Any("error", err))
		return nil, fmt.Errorf("When getting one post,  error:%w", err)
	}
	rows, err := l.db.Query(`
	SELECT 
		c.comment_id,
		c.post_id,
		c.comment_content,
		c.comment_image,
		c.comment_time,
		s.comment_child,
		u1.name AS parent_name,
		u1.avatar_url AS parent_avatar,
		u1.session_id AS parent_session_id,
		sub.post_id AS child_comment_post,
		sub.comment_content AS child_content,
		sub.comment_image AS child_image,
		sub.comment_time AS child_time,
		u2.name AS child_name,
		u2.avatar_url AS child_avatar,
		u2.session_id AS child_session_id
	FROM comments c
	LEFT JOIN subcomments s ON s.comment_parent = c.comment_id
	LEFT JOIN comments sub ON sub.comment_id = s.comment_child
	LEFT JOIN comments_users cu1 ON cu1.comment_id = c.comment_id
	LEFT JOIN users u1 ON u1.user_id = cu1.user_id
	LEFT JOIN comments_users cu2 ON cu2.comment_id = sub.comment_id
	LEFT JOIN users u2 ON u2.user_id = cu2.user_id
	WHERE c.post_id = $1
`, idPost)
	if err != nil {
		log.ErrorContext(ctx, "Error selecting archive_comments", slog.Any("error", err))
		return nil, fmt.Errorf("When selecting archive_comments,  error:%w", err)
	}
	defer rows.Close()

	commentMap := map[int]*archiveCommentNode{}
	for rows.Next() {
		var parentID, parentPostID int
		var parentContent, parentImage, parentName, parentAvatar, parentSession string
		var parentTime time.Time

		var childID, childPostID sql.NullInt64
		var childContent, childImage, childName, childAvatar, childSession sql.NullString
		var childTime sql.NullTime

		err := rows.Scan(
			&parentID, &parentPostID, &parentContent, &parentImage, &parentTime, &childID, // up to index 5
			&parentName, &parentAvatar, &parentSession, // 6–8
			&childPostID, &childContent, &childImage, &childTime, // 9–12
			&childName, &childAvatar, &childSession, // 13–15
		)
		if err != nil {
			log.ErrorContext(ctx, "Error when set of selecting archive_comments to structs", slog.Any("error", err))
			return nil, fmt.Errorf("Error when set of selecting archive_comments to structs, error:%w", err)
		}

		if _, ok := commentMap[parentID]; !ok {
			commentMap[parentID] = &archiveCommentNode{
				parent: archiveOneCommentData{
					id:      parentID,
					postID:  parentPostID,
					content: parentContent,
					image:   parentImage,
					time:    parentTime,
					author: archiveOneCommentAuthor{
						authorName:      parentName,
						authorImageURL:  parentAvatar,
						authorSessionID: parentSession,
					},
				},
			}
		}

		if childID.Int64 != 0 {
			commentMap[parentID].children = append(commentMap[parentID].children, archiveOneCommentData{
				id:      int(childID.Int64),
				postID:  int(childPostID.Int64),
				content: childContent.String,
				image:   childImage.String,
				time:    childTime.Time,
				author: archiveOneCommentAuthor{
					authorName:      childName.String,
					authorImageURL:  childAvatar.String,
					authorSessionID: childSession.String,
				},
			})
		}
	}

	var parentId, childId []int
	for _, value := range commentMap {
		parentId = append(parentId, value.parent.id)
		for _, i := range value.children {
			childId = append(childId, i.id)
		}
	}
	for _, c := range childId {
		for _, p := range parentId {
			if c == p {
				delete(commentMap, c)
			}
		}
	}

	var commentList []archiveCommentNode
	for _, v := range commentMap {
		var childCommentData []archiveOneCommentData
		var parentCommentData archiveOneCommentData
		parentCommentData = v.parent
		for _, child := range v.children {
			if child.id != 0 {
				childCommentData = append(childCommentData, child)
			}
		}
		commentList = append(commentList, archiveCommentNode{
			parent:   parentCommentData,
			children: childCommentData,
		})

	}

	return &archiveOnePostResponse{
		comments: commentList,
		post:     *onePost,
	}, nil
}

func (r *archiveOnePostResponse) GetComments() []database.ArchiveComment {
	comments := make([]database.ArchiveComment, len(r.comments))
	for num, comment := range r.comments {
		comments[num] = &comment
	}
	return comments
}

func (c *archiveCommentNode) GetParent() database.ArchiveOneComment {
	return &c.parent
}

func (c *archiveCommentNode) GetChildren() []database.ArchiveOneComment {
	result := make([]database.ArchiveOneComment, len(c.children))
	for i := range c.children {
		result[i] = &c.children[i]
	}
	return result
}

func (c *archiveOneCommentData) GetCommentID() int                            { return c.id }
func (c *archiveOneCommentData) GetPostID() int                               { return c.postID }
func (c *archiveOneCommentData) GetCommentContent() string                    { return c.content }
func (c *archiveOneCommentData) GetCommentImage() string                      { return c.image }
func (c *archiveOneCommentData) GetCommentTime() time.Time                    { return c.time }
func (c *archiveOneCommentData) GetAuthor() database.ArchiveRespCommentAuthor { return &c.author }

func (o *archiveOneCommentAuthor) GetName() string      { return o.authorName }
func (o *archiveOneCommentAuthor) GetImageURL() string  { return o.authorImageURL }
func (o *archiveOneCommentAuthor) GetSessionID() string { return o.authorSessionID }
