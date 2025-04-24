package controller

import (
	"context"
)

type Interceptor interface {
	InterceptorGov(ctx context.Context) (RespAvatar, error)
}

type Leetboard interface {
	ListPosts(context.Context) (ListPostsResp, error)
	NewPost(context.Context, NewPostReq) (NewPostResp, error)
	ListArchivePosts(context.Context) (ListArchivePostsResp, error)
	OnePostGov(req OnePostReq, ctx context.Context) (OnePostResp, error)
	OneArchivePostGov(ArchiveOnePostReq, context.Context) (ArchiveOnePostResp, error)
	NewComment(NewCommentReq, context.Context) (NewCommentResp, error)
}

type Controller interface {
	Leetboard
	Interceptor
}
