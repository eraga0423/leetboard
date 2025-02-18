package svyazmezhdudbapi

type Forum interface {
	CreatePost()
	ListPosts()
	RemovePost()
	AlterPost()
}

type Svyazmezhdudbapi interface {
	Forum
}
