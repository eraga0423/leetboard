package database

import "context"

type UserAvatars interface {
	ListCharacters(context.Context) (ResponseCharacters, error)
	UpdateCharacters(context.Context, RequestCharacters) error
	InserCartoonCharacters(context.Context, InsertCharacters) error
}

type Post interface {
	ListPosts(context.Context) (ListPostsResp, error)
	CreatePost(context.Context, NewPostReq) (NewPostResp, error)
	OnePost(context.Context, OnePostReq) (OnePostResp, error)
	OneArchivePost(context.Context, ArchiveOnePostReq) (ArchiveOnePostResp, error)
	CreateComment(context.Context, NewReqComment) (NewRespComment, error)
	ListArchivePosts(context.Context) (ListPostsArchiveResp, error)
}

type Database interface {
	Post
	UserAvatars
}
