package posts_handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"text/template"

	"1337b0rd/internal/constants"
)

type reqID struct {
	id int
}

func (h *PostsHandler) GetPostID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tmpl := template.Must(template.ParseFiles(constants.Post))
	id := r.PathValue("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		h.HandleError(w, http.StatusInternalServerError)
		return
	}
	req := new(reqID)
	req = &reqID{id: intID}
	data, err := h.ctrl.OnePostGov(req, ctx)
	if err != nil {
		h.HandleError(w, http.StatusInternalServerError)
		return
	}
	fmt.Println("This GET /post/id")
	err = tmpl.Execute(w, data)
	if err != nil {
		slog.Any("fail", "error")
		return
	}
}
func (i *reqID) GetPostID() int {
	return i.id
}
