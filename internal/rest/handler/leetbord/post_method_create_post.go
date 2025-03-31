package posts_handler

import (
	"fmt"
	"net/http"
)

func (h *PostsHandler) PostMethodCreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This post /create")
}
