package router

import (
	"1337b0rd/internal/rest/handler"
	"context"
	"log"
	"net/http"
)

type Router struct {
	logger  *log.Logger
	router  *http.ServeMux
	handler *handler.Handler
}

func New(h *handler.Handler) *Router {
	mux := http.NewServeMux()
	return &Router{
		router:  mux,
		handler: h,
	}
}

func (r *Router) Start(_ context.Context) *http.ServeMux {
	r.post()
	return r.router
}
