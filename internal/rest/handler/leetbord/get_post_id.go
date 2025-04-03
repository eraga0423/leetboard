package posts_handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"text/template"

	"1337b0rd/internal/constants"
)

func (h *PostsHandler) GetPostID(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(constants.Post))
	id := r.PathValue("id")
	data := id // здесь будет governor
	fmt.Println("This GET /post/id")
	err := tmpl.Execute(w, data)
	if err != nil {
		slog.Any("fail", "error")
		return
	}
}
