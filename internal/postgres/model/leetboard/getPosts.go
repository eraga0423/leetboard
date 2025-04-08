package leetboard

import (
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

type reqPost struct {
	l   *Leetboard
	all allpost
}

func (r reqPost) ListPosts() (database.ListPostsResp, error) {
	rows, err := r.l.db.Query(`
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

	r.all.posts = resPost
	return r.all, nil
}

func (a allpost) GetList() []database.ItemPostsResp {
	var resPosts []database.ItemPostsResp
	for _, v := range a.posts {
		resPosts = append(resPosts, v)
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
