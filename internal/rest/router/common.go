package router

import (
	"context"
	"log"
	"net/http"

	"1337b0rd/internal/rest/handler"
	"1337b0rd/internal/rest/middleware"
)

type Router struct {
	logger  *log.Logger
	router  *http.ServeMux
	handler *handler.Handler
	midd    *middleware.Middleware
}

func New(h *handler.Handler, midd *middleware.Middleware) *Router {
	mux := http.NewServeMux()
	return &Router{
		router:  mux,
		handler: h,
		midd:    midd,
	}
}

func (r *Router) Start(_ context.Context) *http.ServeMux {
	r.post()
	r.style()
	return r.router
}
