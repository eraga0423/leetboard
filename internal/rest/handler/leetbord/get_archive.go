package posts_handler

import (
	"fmt"
	"net/http"
	"text/template"

	constant "1337b0rd/internal/const"
)

func (h *PostsHandler) GetArchive(w http.ResponseWriter, r *http.Request) {
	data := "" //здесь ставить governor 
	fmt.Println("This GET /archive")
	tmpl := template.Must(template.ParseFiles(constant.Archive))
	tmpl.Execute(w, data)
}
