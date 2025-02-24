package posts

import (
	"fmt"
	"net/http"
)

func (h *Posts) GetPosts(w http.ResponseWriter, r *http.Request) {

	fmt.Println("This GET /posts")
}
