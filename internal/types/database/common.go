package database

type UserAvatars interface {
	ListCharacters() (ResponseCharacters, error)
	UpdateCharacters(RequestCharacters) error
	InserCartoonCharacters(InsertCharacters) error
}

type Post interface {
	ListPosts() (ListPostsResp, error)
	CreatePost(NewPostReq) (NewPostResp, error)
	OnePost(OnePostReq) (OnePostResp, error)
	OneArchivePost(ArchiveOnePostReq) (ArchiveOnePostResp, error)
	CreateComment(NewReqComment) error
	ListArchivePosts() (ListPostsArchiveResp, error)
}

type Database interface {
	Post
	UserAvatars
}
