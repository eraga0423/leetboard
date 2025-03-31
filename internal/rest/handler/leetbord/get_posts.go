package posts_handler

import (
	"fmt"
	"net/http"
)

func (h *PostsHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	// Здесь будет 
	fmt.Println("This GET /posts")
}




