package leetboard

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"1337b0rd/internal/types/database"
)

func (l *Leetboard) CreatePost(ctx context.Context, req database.NewPostReq) (database.NewPostResp, error) {
	log := l.logger.With(slog.String("handler", "createPost"))
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
		log.ErrorContext(ctx, "Error starting transaction", slog.Any("error", err))
		return nil, fmt.Errorf("When starting transaction,  error:%w", err)
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
		log.ErrorContext(ctx, "Error inserting data to posts", slog.Any("error", err))
		return nil, fmt.Errorf("When inserting data to posts, error:%w", err)
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
		log.Info("this session_id is first", "name:", authorName)
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
			log.ErrorContext(ctx, "Error inserting data to users", slog.Any("error", err))
			return nil, fmt.Errorf("When inserting data to users, error:%w", err)
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
				log.ErrorContext(ctx, "Error updating data to users", slog.Any("error", err))
				return nil, fmt.Errorf("When updating data to users, error:%w", err)
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
		log.ErrorContext(ctx, "Error inserting data to users_posts", slog.Any("error", err))
		return nil, fmt.Errorf("When inserting data to users_posts, error:%w", err)
	}
	var n newPostRespStruct
	n.newId = int(postID)

	n.tx = tx
	return &n, nil
}

type newPostRespStruct struct {
	newId int
	tx    *sql.Tx
}

func (i *newPostRespStruct) CreationPostResp() int {
	return i.newId
}

func (t newPostRespStruct) TxRollback(b bool) error {
	if b {
		err := t.tx.Rollback()
		if err != nil {
			return err
		}
	} else {
		err := t.tx.Commit()
		if err != nil {
			return err
		}
	}
	return nil
}
