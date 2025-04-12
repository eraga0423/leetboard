package middleware

import (
	"1337b0rd/internal/types/controller"
	"1337b0rd/internal/types/rick_morty"
)

type Middleware struct {
	ctrl      controller.Controller
	mortyRick rick_morty.RestRickAndMorty
}

func New(ctrl controller.Controller) *Middleware {
	return &Middleware{
		ctrl: ctrl,
	}
}
