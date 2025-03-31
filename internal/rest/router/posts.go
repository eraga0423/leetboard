package router

func (r *Router) post() {
	r.router.HandleFunc("GET /catalog", r.handler.GetPosts)
	r.router.HandleFunc("GET /archive", r.handler.GetArchive)
	r.router.HandleFunc("GET /create", r.handler.GetCreatePost)
	r.router.HandleFunc("GET /post/{id}", r.handler.GetPostID)
	r.router.HandleFunc("POST /create", r.handler.PostMethodCreatePost)
}
