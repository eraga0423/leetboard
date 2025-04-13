package posts_handler

import (
	"fmt"
	"net/http"
	"text/template"

	"1337b0rd/internal/constants"
)

func (h *PostsHandler) GetArchive(w http.ResponseWriter, r *http.Request) {
	data := "" // здесь ставить governor
	fmt.Println("This GET /archive")
	tmpl := template.Must(template.ParseFiles(constants.Archive))
	err := tmpl.Execute(w, data)
	if err != nil {
		h.HandleError(w, http.StatusInternalServerError)
		return
	}

}
