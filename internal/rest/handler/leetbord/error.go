package posts_handler

import (
	"log"
	"log/slog"
	"net/http"
	"text/template"

	"1337b0rd/internal/constants"
)

func (h *PostsHandler) HandleError(w http.ResponseWriter, status int) {
	log.Print("method", " handleError")
	errorTmpl := template.Must(template.ParseFiles(constants.Error))
	w.WriteHeader(status)
	err := errorTmpl.Execute(w, map[string]int{
		"Code": status,
	})
	if err != nil {
		slog.Error("method", "handleError", err)
	}
}
