package controller

import "context"

type Interceptor interface {
	InterceptorGov(ctx context.Context) (RespAvatar, error)
	//GetAvatarsInRedis(ctx context.Context) (RespAvatars, error)
}

type Leetboard interface {
	ListPosts(context.Context) (ListPostsResp, error)
	NewPost(context.Context, NewPostReq) (NewPostResp, error)
	// RemovePost(context.Context, string) error
	OnePostGov(req OnePostReq, ctx context.Context) (OnePostResp, error)
}

type Controller interface {
	Leetboard
	Interceptor
}
