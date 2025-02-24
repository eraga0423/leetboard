package handler

import (
	"1337b0rd/internal/governor/posts"
	"1337b0rd/internal/types/controller"
)

type Handler struct {
	*posts.Posts
}

func New(ctrl controller.Controller) *Handler {
	return &Handler{
		Posts: posts.New(ctrl),
	}
}
