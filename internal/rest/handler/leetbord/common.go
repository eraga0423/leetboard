package posts_handler

import (
	"1337b0rd/internal/types/controller"
)

type PostsHandler struct {
	ctrl controller.Controller
	// logger *log.Logger
}

func New(ctrl controller.Controller) *PostsHandler {
	return &PostsHandler{ctrl: ctrl}
}
