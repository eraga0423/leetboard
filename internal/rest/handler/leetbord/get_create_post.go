package posts_handler

import (
	"fmt"
	"net/http"
	"text/template"

	"1337b0rd/internal/constants"
)

func (h *PostsHandler) GetCreatePost(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(constants.CreatePost))
	data := "" // Здесь будет governor
	tmpl.Execute(w, data)
	fmt.Println("This GET /create")
}
