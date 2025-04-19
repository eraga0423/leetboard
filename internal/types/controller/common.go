package controller

import (
	"context"
)

type Interceptor interface {
	InterceptorGov(ctx context.Context) (RespAvatar, error)
	//GetAvatarsInRedis(ctx context.Context) (RespAvatars, error)
}

type Leetboard interface {
	ListPosts(context.Context) (ListPostsResp, error)
	NewPost(context.Context, NewPostReq) (NewPostResp, error)
	ListArchivePosts(context.Context) (ListArchivePostsResp, error)
	OnePostGov(req OnePostReq, ctx context.Context) (OnePostResp, error)
	OneArchivePostGov(ArchiveOnePostReq, ctx context.Context) (ArchiveOnePostResp, error)
}

type Controller interface {
	Leetboard
	Interceptor
}
