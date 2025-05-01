package leetboard

import (
	"database/sql"
	"errors"
	"log/slog"
	"time"

	"1337b0rd/internal/types/database"
)

func (a *Leetboard) CreateComment(req database.NewReqComment) (database.NewRespComment, error) {
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
		return nil, err
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
		return nil, err
	}
	if parentComID != 0 {
		_, err := tx.Exec(`
	INSERT INTO subcomments
	(comment_parent, comment_child)
	VALUE($1, $2)
	`, parentComID, commentID,
		)
		if err != nil {
			return nil, err
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
	} else if err != nil {
		return nil, err
	}

	_, err = tx.Exec(`
	INSERT INTO comments_users
	(comment_id, user_id)
	VALUE($1, $2)
	`, commentID, authorID,
	)
	if err != nil {
		return nil, err
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
