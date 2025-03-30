package controller

import "context"

//type Interseptor interface {
//	Authentificator(context.Context, string)(context.Context, error)
//}

type Auth interface {
	SignIn(ctx context.Context) (SignInResp, error)
	SignUp(ctx context.Context) (SignUpResp, error)
}

type Leetboard interface {
	ListPosts(context.Context) (ListPostsResp, error)
	NewPost(context.Context, NewPostReq) (NewPostResp, error)
	RemovePost(context.Context, string) error
}

type Controller interface {
	Leetboard
	Auth
}
