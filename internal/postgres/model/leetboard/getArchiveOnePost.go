package leetboard

import (
	"1337b0rd/internal/types/database"
	"database/sql"
	"time"
)

type archiveOnePost struct {
	onePostTitle    string
	onePostContent  string
	onePostURLImage string
	onePostTime     time.Time
	author          archiveOnePostAuthor
}

type archiveOnePostAuthor struct {
	authorName      string
	authorImageURL  string
	authorSessionID string
}

func archiveReturnOnePost(idPost int, db *sql.DB) (archiveOnePost, error) {
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
	WHERE 
	p.post_id = $1
	AND p.deletion = TRUE`, idPost)
	var o archiveOnePost
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
		return archiveOnePost{}, err
	}
	return o, nil
}
func (r archiveOnePostResponse) GetOnePost() database.ArchiveRespOnePost  { return r.post }
func (o archiveOnePost) GetTitle() string                                 { return o.onePostTitle }
func (o archiveOnePost) GetPostContent() string                           { return o.onePostContent }
func (o archiveOnePost) GetPostUrlImage() string                          { return o.onePostURLImage }
func (o archiveOnePost) GetPostTime() time.Time                           { return o.onePostTime }
func (o archiveOnePost) GetAuthorPost() database.ArchiveRespOnePostAuthor { return o.author }

func (o archiveOnePostAuthor) GetName() string      { return o.authorName }
func (o archiveOnePostAuthor) GetImageURL() string  { return o.authorImageURL }
func (o archiveOnePostAuthor) GetSessionID() string { return o.authorSessionID }
