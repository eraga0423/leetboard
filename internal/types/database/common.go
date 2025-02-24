package database

type Post interface {
	ListPosts()
	CreatePost()
	UpdatePost()
	DeletePost()
}

type Database interface {
	Post
}
