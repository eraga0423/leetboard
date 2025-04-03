package posts_handler

import (
	"fmt"
	"net/http"
	"text/template"

	"1337b0rd/internal/constants"
)

func (h *PostsHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This GET /posts")
	data := "" // Здесь будет governor
	tmpl := template.Must(template.ParseFiles(constants.Catalog))

	tmpl.Execute(w, data)
}
