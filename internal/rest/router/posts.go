package router

import (
	"net/http"

	"1337b0rd/internal/constants"
)

func (r *Router) post() {
	r.router.Handle("GET /catalog", r.midd.Authentificator(http.HandlerFunc(r.handler.GetPosts)))
	r.router.Handle("GET /archive", r.midd.Authentificator(http.HandlerFunc(r.handler.GetArchive)))
	r.router.Handle("GET /create", r.midd.Authentificator(http.HandlerFunc(r.handler.GetCreatePost)))
	r.router.Handle("GET /post/{id}", r.midd.Authentificator(http.HandlerFunc(r.handler.GetPostID)))
	r.router.Handle("POST /create", r.midd.Authentificator(http.HandlerFunc(r.handler.PostMethodCreatePost)))
	r.router.Handle("POST /post/{id}", r.midd.Authentificator(http.HandlerFunc(r.handler.NewComment)))
	r.router.Handle("GET /archive/{id}", r.midd.Authentificator(http.HandlerFunc(r.handler.GetPostIDArchive)))
}

func (r *Router) style() {
	fs := http.FileServer(http.Dir(constants.DirCss))
	r.router.Handle("/static/", http.StripPrefix("/static/", fs))
}
