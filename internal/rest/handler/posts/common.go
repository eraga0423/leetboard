package posts

import (
	"1337b0rd/internal/types/controller"
	"log"
)

type Posts struct {
	ctrl   controller.Controller
	logger *log.Logger
}

func New(ctrl controller.Controller, logger *log.Logger) *Posts {
	return &Posts{ctrl: ctrl, logger: logger}
}
