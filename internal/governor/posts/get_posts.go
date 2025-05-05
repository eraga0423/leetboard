package posts_governor

import (
	"context"
	"log"
	"time"

	"1337b0rd/internal/types/controller"
)

func (r *PostsGovernor) ListPosts(ctx context.Context) (controller.ListPostsResp, error) {
	var newResp []postResp
	dataBaseResp, err := r.db.ListPosts(ctx)
	if err != nil {
		log.Print("method:", " listsPost", " err:", err, " dir:", " governor")
		return nil, err
	}
	listResp := dataBaseResp.GetList()
	for _, v := range listResp {
		newResp = append(newResp, postResp{
			PostID:      v.GetPostID(),
			PostTitle:   v.GetTitle(),
			PostContent: v.GetPostContent(),
			PostImage:   v.GetPostImageURL(),
			PostTime:    v.GetPostTime(),
		})
	}

	r.all.posts = newResp
	return &r.all, nil
}

type allPost struct {
	posts []postResp
}
type postResp struct {
	PostID      int
	PostTitle   string
	PostContent string
	PostImage   string
	PostTime    time.Time
}

func (p *allPost) GetList() []controller.ItemPostsResp {
	res := make([]controller.ItemPostsResp, len(p.posts))
	for i, v := range p.posts {
		res[i] = &v
	}
	return res
}

func (p *postResp) GetPostID() int {
	return p.PostID
}

func (p *postResp) GetTitle() string {
	return p.PostTitle
}

func (p *postResp) GetPostContent() string {
	return p.PostContent
}

func (p *postResp) GetPostImageURL() string {
	return p.PostImage
}

func (p *postResp) GetPostTime() time.Time {
	return p.PostTime
}
