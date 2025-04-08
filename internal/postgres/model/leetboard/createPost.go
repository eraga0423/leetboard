package leetboard

import (
	"1337b0rd/internal/types/database"
	"errors"
	"time"
)

func (l *Leetboard) CreatePost(req database.NewPostReq) (database.NewPostResp, error) {
	title := req.GetTitle()
	content := req.GetPostContent()
	image := req.GetImage()
	authID := req.GetAuthor()
	curTime := time.Now()
	sql := l.db.QueryRow(`
	INSERT INTO posts
	(title, post_content, post_image, post_time)
	VALUE($1, $2, $3, $4)
	RETURING post_id
	`, title, content, image, curTime,
	)
	var postID int
	err := sql.Scan(
		&postID,
	)
	if err != nil {
		return nil, err
	}
	if postID == 0 {
		return nil, errors.New("idpost does not assign")
	}
	_, err = l.db.Exec(`
	INSERT INTO user_post
	(post_id, user_id)
	VALUE($1, $2)
	`, postID, authID,
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
