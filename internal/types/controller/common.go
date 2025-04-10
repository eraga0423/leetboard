package controller

import "context"

type Interceptor interface {
	InterceptorGov(ctx context.Context, sessionID string) context.Context
	GenerateSessionID() (string, error)
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
