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
	logger  *slog.Logger
	router  *router.Router
	handler *handler.Handler
}

func New(ctrl controller.Controller, logger *slog.Logger) *Rest {
	h := handler.New(ctrl)
	m := middleware.New(ctrl)
	r := router.New(h, m)

	return &Rest{
		logger:  logger,
		handler: h,
		router:  r,
	}
}

func (r *Rest) Start(ctx context.Context, cancelFunc context.CancelFunc, conf *config.APIConfig, _ *slog.Logger) error {
	err := mime.AddExtensionType(".css", "text/css")
	if err != nil {
		r.logger.ErrorContext(ctx, "add extension type css error", slog.Any("error", err))
		return fmt.Errorf("Error when add extension type css: %w", err)
	}
	mux := r.router.Start(ctx)
	srv := &http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%s", conf.Rest.Port),
	}

	log.Print("server start port: ", conf.Rest.Port)
	go func(cancelFunc context.CancelFunc) {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			r.logger.ErrorContext(ctx, "Error when start listenandserve", slog.Any("error", err))
			cancelFunc()
		}
	}(cancelFunc)

	<-ctx.Done()
	if err := srv.Shutdown(ctx); err != nil {
		r.logger.ErrorContext(ctx, "Error when close server", slog.Any("error", err))
		return fmt.Errorf("Error when close server: %w", err)
	} else {
		r.logger.Info("Server succesfully closed")
	}

	return nil
}
