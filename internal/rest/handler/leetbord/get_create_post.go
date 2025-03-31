package posts_handler

import (
	"fmt"
	"net/http"
)

func (h *PostsHandler) GetCreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This GET /create")
}
