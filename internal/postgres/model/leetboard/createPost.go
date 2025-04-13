package leetboard

import (
	"time"

	"1337b0rd/internal/types/database"
)

func (l *Leetboard) CreatePost(req database.NewPostReq) (database.NewPostResp, error) {
	title := req.GetTitle()
	content := req.GetPostContent()
	image := req.GetImage()
	authSessionID := req.GetAuthorSession()
	curTime := time.Now()

	tx, err := l.db.Begin()
	if err != nil {
		return nil, err
	}
	defer TxAfter(tx, err)

	sql := tx.QueryRow(`
	INSERT INTO posts
	(title, post_content, post_image, post_time)
	VALUE($1, $2, $3, $4)
	RETURING post_id
	`, title, content, image, curTime,
	)
	var postID, authorID int
	err = sql.Scan(
		&postID,
	)
	if err != nil {
		return nil, err
	}

	sql = tx.QueryRow(`
	SELECT user_id
	FROM users
	WHERE session_id = $1
	`, authSessionID)
	err = sql.Scan(
		&authorID,
	)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(`
	INSERT INTO user_post
	(post_id, user_id)
	VALUE($1, $2)
	`, postID, authorID,
	)
	if err != nil {
		return nil, err
	}
	return idPost{newId: int(postID)}, nil
}

type idPost struct {
	newId int
}

func (i idPost) CreationPostResp() int {
	return i.newId
}
