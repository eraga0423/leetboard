package posts_handler

import (
	"fmt"
	"net/http"
	"text/template"

	"1337b0rd/internal/constants"
)

func (h *PostsHandler) GetError(w http.ResponseWriter, r *http.Request) {
	data := "" // здесь ставить governor
	fmt.Println("This GET /error")
	tmpl := template.Must(template.ParseFiles(constants.Error))
	tmpl.Execute(w, data)
}
