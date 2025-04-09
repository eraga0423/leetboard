package database

type User interface {
	FindUser(FindUserReq) (bool, FindUserResp)
}

type Post interface {
	ListPosts() (ListPostsResp, error)
	// CreatePost(NewPostReq) (NewPostResp, error)
	// DeletePost(RemovePostReq) (bool, error)
	// OnePost(OnePostReq) (OnePostResp, error)
	// CreateComment(idPost string) error
}

type Database interface {
	Post
	User
}
