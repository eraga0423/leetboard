package leetboard

import "1337b0rd/internal/types/database"

func (l *Leetboard) DeletePost(r database.RemovePostReq) (bool, error) {
	idPost := r.GetPostID()

	_, err := l.db.Exec(`
	UPDATE posts
	deletion = TRUE
	WHERE post_id = $1
	`, idPost)
	if err != nil {
		return false, err
	}
	return true, nil

}
