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

func (a *Leetboard) CreateComment(ctx context.Context, req database.NewReqComment) (database.NewRespComment, error) {
	log := a.logger.With(slog.String("handler", "CreateComment"))

	// comment
	post := req.GetPostID()
	content := req.GetCommentContent()
	comImage := req.GetCommentImage()

	// parent comment
	parentComID := req.GetParentCommentID()

	// comment's author
	authSessionID := req.GetAuthorSession()
	authorName := req.GetAuthorName()
	authorAvatar := req.GetAuthorAvatarURL()

	// time
	timeNow := time.Now()

	tx, err := a.db.Begin()
	if err != nil {
		log.ErrorContext(ctx, "Begin starts a transaction error", slog.Any("error", err))
		return nil, fmt.Errorf("Error when start transaction: %w", err)
	}
	defer TxAfter(tx, err)

	mydb := tx.QueryRow(`
	INSERT INTO comments
	(post_id, comment_content, comment_image, comment_time)
	VALUE($1, $2, $3, $4)
	RETURING comment_id
	`, post, content, comImage, timeNow,
	)
	var commentID int
	err = mydb.Scan(
		&commentID,
	)
	if err != nil {
		log.ErrorContext(ctx, "insert comment's error", slog.Any("error", err))
		return nil, fmt.Errorf("Error when insert comments: %w", err)
	}
	if parentComID != 0 {
		_, err := tx.Exec(`
	INSERT INTO subcomments
	(comment_parent, comment_child)
	VALUE($1, $2)
	`, parentComID, commentID,
		)
		if err != nil {
			log.ErrorContext(ctx, "insert subcomment's error", slog.Any("error", err))
			return nil, fmt.Errorf("Error when insert subcomments: %w", err)
		}
	}
	var authorID int
	mydb = tx.QueryRow(`
	SELECT user_id
	FROM users
	WHERE session_id = $1
	`, authSessionID)
	err = mydb.Scan(
		&authorID,
	)
	if errors.Is(err, sql.ErrNoRows) {
		log.Info("this session_id is first", "author name:", authorName)
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
			log.ErrorContext(ctx, "insert user error", slog.Any("error", err))
			return nil, fmt.Errorf("Error when insert user: %w", err)
		}
	} else if err != nil {
		log.ErrorContext(ctx, "select user id error", slog.Any("error", err))
		return nil, fmt.Errorf("Error when select user id: %w", err)
	}

	_, err = tx.Exec(`
	INSERT INTO comments_users
	(comment_id, user_id)
	VALUE($1, $2)
	`, commentID, authorID,
	)
	if err != nil {
		log.ErrorContext(ctx, "insert comments_users error", slog.Any("error", err))
		return nil, fmt.Errorf("Error when insert comments_users: %w", err)
	}
	var d txRollBackStruct
	d.tx = tx

	return &d, nil
}

type txRollBackStruct struct {
	tx *sql.Tx
}

func (t *txRollBackStruct) TxRollback(b bool) error {
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
