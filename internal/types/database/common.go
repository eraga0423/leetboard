package database

import "context"

type User interface {
	FindUser(context.Context, FindUserReq) (bool, FindUserResp)
}

type Post interface {
	ListPosts() (ListPostsResp, error)
	CreatePost(NewPostReq) (NewPostResp, error)
	DeletePost(RemovePostReq) (bool, error)
	OnePost(OnePostReq) (OnePostResp, error)
	// CreateComment(idPost string) error
}

type Database interface {
	Post
	User
}
