package router

import "net/http"

func (r *Router) post() {
	r.router.Handle("GET /catalog", r.midd.Authentificator(http.HandlerFunc(r.handler.GetPosts)))
	r.router.Handle("GET /archive", r.midd.Authentificator(http.HandlerFunc(r.handler.GetArchive)))
	r.router.Handle("GET /create", r.midd.Authentificator(http.HandlerFunc(r.handler.GetCreatePost)))
	r.router.Handle("GET /post/{id}", r.midd.Authentificator(http.HandlerFunc(r.handler.GetPostID)))
	r.router.Handle("POST /create", r.midd.Authentificator(http.HandlerFunc(r.handler.PostMethodCreatePost)))
	r.router.Handle("GET /error", r.midd.Authentificator(http.HandlerFunc(r.handler.GetError)))
}
