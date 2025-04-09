package posts_handler

import (
	"1337b0rd/internal/constants"
	"log"
	"net/http"
	"text/template"
)

func (h *PostsHandler) HandleError(w http.ResponseWriter, status int) {
	var errorTmpl = template.Must(template.ParseFiles(constants.Error))
	w.WriteHeader(status)
	err := errorTmpl.Execute(w, map[string]int{
		"Code": status,
	})
	if err != nil {
		log.Print("method", "handleError")
		log.Print(err)
	}
}
