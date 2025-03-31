package posts_governor

import (
	"1337b0rd/internal/types/controller"
)

type PostsGovernor struct {
	// logger *log.Logger
	ctrl controller.Controller
}

func New(ctrl controller.Controller) *PostsGovernor {
	return &PostsGovernor{
		// logger: logger,
		ctrl: ctrl,
	}
}
