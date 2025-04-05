package leetboard

import (
	"1337b0rd/internal/types/database"
	"time"
)

type postResp struct {
	postID      int
	postTitle   string
	postContent string
	postImage   string
	postTime    time.Time
}

func (l Leetboard) GetList() ([]database.ItemPostsResp, error) {
	rows, err := l.db.Query(`
	SELECT 
	post_id,
	title,
	post_content,
	post_image,
	post_time
	 FROM posts`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var resPost []database.ItemPostsResp
	for rows.Next() {
		var p postResp
		err := rows.Scan(
			&p.postID,
			&p.postTitle,
			&p.postContent,
			&p.postImage,
			&p.postTime,
		)
		if err != nil {
			return nil, err
		}
		resPost = append(resPost, p)
	}
	return resPost, nil

}
func (p postResp) GetPostID() int {
	return p.postID
}
func (p postResp) GetTitle() string {
	return p.postTitle
}
func (p postResp) GetPostContent() string {
	return p.postContent
}
func (p postResp) GetPostImageURL() string {
	return p.postImage
}
func (p postResp) GetPostTime() time.Time {
	return p.postTime
}
