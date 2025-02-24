package posts

import (
	"1337b0rd/internal/types/controller"
)

type Posts struct {
	//logger *log.Logger
	ctrl controller.Controller
}

func New(ctrl controller.Controller) *Posts {
	return &Posts{
		//logger: logger,
		ctrl: ctrl,
	}
}
