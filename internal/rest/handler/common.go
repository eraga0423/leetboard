package handler

import (
	posts_handler "1337b0rd/internal/rest/handler/leetbord"
	"1337b0rd/internal/types/controller"
)

type Handler struct {
	*posts_handler.PostsHandler
}

func New(ctrl controller.Controller) *Handler {
	return &Handler{
		PostsHandler: posts_handler.New(ctrl),
	}
}
