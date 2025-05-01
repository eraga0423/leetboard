package rest

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"mime"
	"net/http"

	"1337b0rd/internal/config"

	"1337b0rd/internal/rest/handler"
	"1337b0rd/internal/rest/middleware"
	"1337b0rd/internal/rest/router"
	"1337b0rd/internal/types/controller"
)

type Rest struct {
	// logger *log.Logger
	router  *router.Router
	handler *handler.Handler
}

func New(ctrl controller.Controller) *Rest {
	h := handler.New(ctrl)
	m := middleware.New(ctrl)
	r := router.New(h, m)

	return &Rest{
		// logger: logger,
		handler: h,
		router:  r,
	}
}

func (r *Rest) Start(ctx context.Context, cancelFunc context.CancelFunc, conf *config.APIConfig, logger *slog.Logger) error {
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
	go func(cancelFunc context.CancelFunc) {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Print("add extension type css error", err.Error())
		}
		cancelFunc()
	}(cancelFunc)

	<-ctx.Done()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Ошибка при завершении работы сервера: %v\n", err)
	} else {
		fmt.Println("Сервер успешно завершен.")
	}

	return nil
}
