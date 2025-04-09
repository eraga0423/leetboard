package rest

import (
	"1337b0rd/internal/config"
	"context"
	"fmt"
	"log"
	"mime"
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

func (r *Rest) Start(ctx context.Context, conf *config.APIConfig) error {
	err := mime.AddExtensionType(".css", "text/css")
	if err != nil {
		log.Print("add extension type css error")
		return err
	}
	mux := r.router.Start(ctx)
	srv := &http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%s", conf.Rest.Port),
	}
	log.Print("server start port: ", conf.Rest.Port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Print("add extension type css error", err.Error())
		return err
	}
	return nil
}
