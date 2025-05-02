package posts_handler

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"text/template"

	"1337b0rd/internal/constants"
)

type resp struct {
	TitlePost string
}

func (h *PostsHandler) GetCreatePost(w http.ResponseWriter, r *http.Request) {
	slog.Info("get create post")
	tmpl := template.Must(template.ParseFiles(constants.CreatePost))
	data := resp{
		TitlePost: "NEW POST",
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
		h.HandleError(w, http.StatusBadRequest)
		return
	}
	fmt.Println("This GET /create")
}
