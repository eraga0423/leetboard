package leetboard

import (
	"time"

	"1337b0rd/internal/types/database"
)

func (a *Leetboard) CreateComment(req database.NewReqComment) error {
	post := req.GetPostID()
	parentComID := req.GetParentCommentID()
	content := req.GetCommentContent()
	comImage := req.GetCommentImage()
	authSessionID := req.GetAuthorSession()
	timeNow := time.Now()

	tx, err := a.db.Begin()
	if err != nil {
		return err
	}
	defer TxAfter(tx, err)

	sql := tx.QueryRow(`
	INSERT INTO comments
	(post_id, comment_content, comment_image, comment_time)
	VALUE($1, $2, $3, $4)
	RETURING comment_id
	`, post, content, comImage, timeNow,
	)
	var commentID int
	err = sql.Scan(
		&commentID,
	)
	if err != nil {
		return err
	}
	if parentComID != 0 {
		_, err := tx.Exec(`
	INSERT INTO subcomments
	(comment_parent, comment_child)
	VALUE($1, $2)
	`, parentComID, commentID,
		)
		if err != nil {
			return err
		}
	}
	var authorID int
	sql = tx.QueryRow(`
	SELECT user_id
	FROM users
	WHERE session_id = $1
	`, authSessionID)
	err = sql.Scan(
		&authorID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
	INSERT INTO comments_users
	(comment_id, user_id)
	VALUE($1, $2)
	`, commentID, authorID,
	)
	if err != nil {
		return err
	}
	return nil
}
