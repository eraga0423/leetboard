package database

type Post interface {
	ListPosts() error
	CreatePost(idPost string) error
	UpdatePost(idPost string) error
	DeletePost() error
	CreateComment(idPost string) error
}

type Database interface {
	Post
}
