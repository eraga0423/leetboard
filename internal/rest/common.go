package rest

import (
	"context"
	"net/http"

	"1337b0rd/internal/rest/handler"
	"1337b0rd/internal/rest/middleware"
	"1337b0rd/internal/rest/router"
	"1337b0rd/internal/types/controller"
)

type Rest struct {
	// logger *log.Logger
	router *router.Router
}

func New(ctrl controller.Controller) *Rest {
	h := handler.New(ctrl)
	m := middleware.New(ctrl)
	r := router.New(h, m)

	return &Rest{
		// logger: logger,
		router: r,
	}
}

func (r *Rest) Start(ctx context.Context) error {
	mux := r.router.Start(ctx)
	srv := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
