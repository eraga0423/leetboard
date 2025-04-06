package database

type Post interface {
	ListPosts() (ListPostsResp, error)
	CreatePost(NewPostReq) (NewPostResp, error)
	DeletePost(RemovePostReq) error
	OnePost(OnePostReq) (OnePostResp, error)
	// CreateComment(idPost string) error
}

type Database interface {
	Post
}
