package middleware

import "1337b0rd/internal/types/controller"

type Middleware struct {
	ctrl controller.Controller
}

func New(ctrl controller.Controller) *Middleware {
	return &Middleware{
		ctrl: ctrl,
	}
}
