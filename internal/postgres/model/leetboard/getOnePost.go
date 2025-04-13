package leetboard

import (
	"database/sql"
	"errors"
	"time"

	"1337b0rd/internal/types/database"
)

type onePost struct {
	onePostTitle    string
	onePostContent  string
	onePostURLImage string
	onePostTime     time.Time
	author          onePostAuthor
}

type onePostAuthor struct {
	authorName      string
	authorImageURL  string
	authorSessionID string
}

func returnOnePost(idPost int, db *sql.DB) (onePost, error) {
	sql := db.QueryRow(`
	SELECT 
	p.title, 
	p.post_content, 
	p.post_image, 
	p.post_time,
	u.name,
	u.user_avatar,
	u.session_id
	FROM posts p
	LEFT JOIN users_posts up ON up.post_id=p.post_id
	LEFT JOIN users u ON u.user_id = up.user_id
	LEFT JOIN comments c ON c.post_id = p.post_id
	WHERE (
    (
        c.comment_time IS NOT NULL AND 
        c.comment_time >= NOW() - INTERVAL '15 minutes'
    ) 
    OR 
    (
        c.comment_time IS NULL AND 
        p.post_time >= NOW() - INTERVAL '10 minutes'
    )
)
AND p.post_id = $1`, idPost)
	var o onePost
	err := sql.Scan(
		&o.onePostTitle,
		&o.onePostContent,
		&o.onePostURLImage,
		&o.onePostTime,
		&o.author.authorName,
		&o.author.authorImageURL,
		&o.author.authorSessionID,
	)
	if err != nil {
		return onePost{}, err
	}
	if o.onePostTitle == "" {
		return onePost{}, errors.New("post empty")
	}
	return o, nil
}
func (r onePostResponse) GetOnePost() database.RespOnePost  { return r.post }
func (o onePost) GetTitle() string                          { return o.onePostTitle }
func (o onePost) GetPostContent() string                    { return o.onePostContent }
func (o onePost) GetPostUrlImage() string                   { return o.onePostURLImage }
func (o onePost) GetPostTime() time.Time                    { return o.onePostTime }
func (o onePost) GetAuthorPost() database.RespOnePostAuthor { return o.author }

func (o onePostAuthor) GetName() string      { return o.authorName }
func (o onePostAuthor) GetImageURL() string  { return o.authorImageURL }
func (o onePostAuthor) GetSessionID() string { return o.authorSessionID }
