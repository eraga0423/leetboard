package posts_governor

import (
	"context"
	"time"

	"1337b0rd/internal/types/controller"
)

func (r *PostsGovernor) ListPosts(ctx context.Context) (controller.ListPostsResp, error) {

	// sessionID, ok := ctx.Value(constants.SessionIDKey).(string)
	// if !ok {
	// 	return nil, fmt.Errorf("user not found in context")
	// }
	// ps := ma
	return nil, nil
}

type ListPlayerResp struct {
	Posts []PostResp
}
type PostResp struct {
	PostID           int
	PostTitle        string
	PostContent      string
	PostImage        string
	PostTime         time.Time
	UserResp         User
	CommentsResponse []MonoCommentResp
}

type User struct {
	UserID     int
	UserName   string
	UserAvatar string
}

type MonoCommentResp struct {
	CommentID      int
	PostID         int
	CommentContent string
	CommentImage   string
	CommentTime    time.Time
	SubComments    []MonoCommentResp
}

func (i *ListPlayerResp) GetList() []controller.ItemPostsResp {
	list := make([]controller.ItemPostsResp, 0, len(i.Posts))
	for _, p := range i.Posts {
		list = append(list, p)
	}
	return list
}

func newPost(
	postID int,
	title,
	content,
	imageURL string, postTime time.Time,
	user User, comments CommentsResp,
) *PostResp {
	return &PostResp{
		PostID:       postID,
		PostTitle:    title,
		PostContent:  content,
		PostImage:    imageURL,
		PostTime:     postTime,
		User:         user,
		CommentsResp: comments,
	}
}

func (p *PostResp) GetPostID() int {
	return p.PostID
}

func (p *PostResp) GetTitle() string {
	return p.PostTitle
}

func (p *PostResp) GetPostContent() string {
	return p.PostContent
}

func (p *PostResp) GetPostImageURL() string {
	return p.PostImage
}

func (p *PostResp) GetPostTime() time.Time {
	return p.PostTime
}

func (p *PostResp) GetUser() User {
	return p.User
}

func (p *PostResp) GetComments() []CommentResp {
	return p.CommentsResp
}

func (p *PostResp) GetUserID() int {
	return p.UserID
}

func (p *PostResp) GetName() string {
	return p.UserName
}

func (p *PostResp) GetAvatar() string {
	return p.UserAvatar
}

func (p *CommentsResp) GetParentComment() []MonoCommentResp {
	list := make([]controller.ItemMonoComment, 0, len(p.CommentsResp))
	for _, c := range p.CommentsResp {
		list = append(list, c.ParentComment)
	}
	return list
}

func (p *CommentsResp) GetSubComment() []MonoCommentResp {
	list := make([]controller.ItemMonoComment, 0)
	for _, c := range p {
		list = append(list, c)
	}
	return list
}

func (p *MonoCommentResp) GetCommentID() int {
	return p.CommentID
}
