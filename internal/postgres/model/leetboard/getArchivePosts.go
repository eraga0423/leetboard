package leetboard

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"1337b0rd/internal/types/database"
)

type allArchivepost struct {
	posts []postArchiveResp
}

type postArchiveResp struct {
	postID      int
	postTitle   string
	postContent string
	postImage   string
	postTime    time.Time
}

func (l *Leetboard) ListArchivePosts(ctx context.Context) (database.ListPostsArchiveResp, error) {
	log := l.logger.With(slog.String("handler", "ListArchivePost"))

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
    c.comment_time < NOW() - INTERVAL '15 minutes'
) OR (
   
    c.comment_time IS NULL AND 
    p.post_time < NOW() - INTERVAL '10 minutes'
)`)
	if err != nil {
		log.ErrorContext(ctx, "Error selecting archive_posts", slog.Any("error", err))
		return nil, fmt.Errorf("Error when selecting archive_posts: %w", err)
	}
	defer rows.Close()

	resPost := make(map[int]postArchiveResp)
	var responseArchivePost []postArchiveResp
	for rows.Next() {
		var p postArchiveResp
		err := rows.Scan(
			&p.postID,

			&p.postTitle,
			&p.postContent,
			&p.postImage,
			&p.postTime,
		)
		if err != nil {
			log.ErrorContext(ctx, "Error set archive_posts of selecting to structs", slog.Any("error", err))
			return nil, fmt.Errorf("Error set archive_posts of selecting to structs: %w", err)
		}
		val, exist := resPost[p.postID]
		if !exist {
			responseArchivePost = append(responseArchivePost, p)
			resPost[p.postID] = val
		}

	}

	return &allArchivepost{posts: responseArchivePost}, nil
}

func (a *allArchivepost) GetArchiveList() []database.ItemPostsArchiveResp {
	resPosts := make([]database.ItemPostsArchiveResp, len(a.posts))
	for num, post := range a.posts {
		resPosts[num] = &post
	}
	return resPosts
}

func (p *postArchiveResp) GetPostID() int {
	return p.postID
}

func (p *postArchiveResp) GetTitle() string {
	return p.postTitle
}

func (p *postArchiveResp) GetPostContent() string {
	return p.postContent
}

func (p *postArchiveResp) GetPostImageURL() string {
	return p.postImage
}

func (p *postArchiveResp) GetPostTime() time.Time {
	return p.postTime
}
