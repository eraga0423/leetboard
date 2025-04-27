package leetboard

import (
	"fmt"
	"log/slog"
	"time"

	"1337b0rd/internal/types/database"
)

type allpost struct {
	posts []postResp
}

type postResp struct {
	postID      int
	postTitle   string
	postContent string
	postImage   string
	postTime    time.Time
}

func (l *Leetboard) ListPosts() (database.ListPostsResp, error) {
	slog.Info("start get getposts")
	rows, err := l.db.Query(`
    SELECT 
    p.post_id,
    p.title,
    p.post_content,
    p.post_image,
    p.post_time
FROM posts p
LEFT JOIN comments c ON c.post_id = p.post_id
WHERE (
   
    c.comment_time IS NOT NULL AND 
    c.comment_time >= NOW() - INTERVAL '15 minutes'
) OR (
   
    c.comment_time IS NULL AND 
    p.post_time >= NOW() - INTERVAL '10 minutes'
)
`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var resPost []postResp
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
	slog.Info("end get posts")
	return allpost{posts: resPost}, nil
}

func (a allpost) GetList() []database.ItemPostsResp {
	resPosts := make([]database.ItemPostsResp, len(a.posts))
	for num, post := range a.posts {
		resPosts[num] = post
	}
	return resPosts
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
