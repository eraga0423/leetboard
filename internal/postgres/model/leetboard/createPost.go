package leetboard

import (
	"database/sql"
	"errors"
	"log/slog"
	"time"

	"1337b0rd/internal/types/database"
)

func (l *Leetboard) CreatePost(req database.NewPostReq) (database.NewPostResp, error) {
	// post
	title := req.GetTitle()
	content := req.GetPostContent()
	image := req.GetPostImageURL()

	// user
	authorAvatar := req.GetAuthorAvatarURL()
	authorName := req.GetAuthorName()
	authSessionID := req.GetAuthorSession()

	// time
	curTime := time.Now()

	// start
	tx, err := l.db.Begin()
	if err != nil {
		return nil, err
	}
	defer TxAfter(tx, err)

	// post script
	mydb := tx.QueryRow(`
	INSERT INTO posts
	(title, post_content, post_image, post_time)
	VALUES($1, $2, $3, $4)
	RETURNING post_id
	`, title, content, image, curTime,
	)
	var postID, authorID int
	err = mydb.Scan(
		&postID,
	)
	if err != nil {
		return nil, err
	}

	// user script
	mydb = tx.QueryRow(`
	SELECT user_id, name
	FROM users
	WHERE session_id = $1`, authSessionID)
	var oldName string

	err = mydb.Scan(
		&authorID,
		&oldName,
	)

	if errors.Is(err, sql.ErrNoRows) {
		slog.Info("this session_id is first")
		// new user insert
		mydb = tx.QueryRow(`
		INSERT INTO users 
		(name, avatar_url, session_id)
		VALUES($1, $2, $3)
		RETURNING user_id
		`, authorName, authorAvatar, authSessionID)
		err = mydb.Scan(
			&authorID,
		)
		if err != nil {
			return nil, err
		}

	} else {
		if oldName != authorName {
			slog.Info("update user name")
			// update user name
			_, err = tx.Exec(`
			UPDATE users 
			SET name = $1
			WHERE session_id = $2
			`, authorName, authSessionID)
			if err != nil {
				return nil, err
			}
		}
	}

	_, err = tx.Exec(`
	INSERT INTO users_posts
	(post_id, user_id)
	VALUES($1, $2)
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
