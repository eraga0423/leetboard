package posts_handler

import (
	"fmt"
	"net/http"
)

func (h *PostsHandler) GetPostID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This GET /post/id")
}
