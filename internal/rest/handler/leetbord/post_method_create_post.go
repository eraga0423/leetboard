package posts_handler

import (
	"1337b0rd/internal/constants"
	"fmt"
	"io"
	"log"
	"net/http"
)

type req struct {
	authorIDSession string
	title           string
	content         string
	nick            string
	fileIO          io.Reader
}

func (h *PostsHandler) PostMethodCreatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "file error", http.StatusBadRequest)
		return
	}
	fmt.Println("This post /create")
	NewReq := new(req)
	name := r.FormValue("name")
	title := r.FormValue("title")
	postText := r.FormValue("post")
	file, _, err := r.FormFile("image")
	authorID, err := r.Cookie(constants.SessionIDKey)

	if err != nil && err != http.ErrMissingFile {
		http.Error(w, "Не удалось добавить файл", http.StatusBadRequest)
		return
	}

	defer file.Close()

	NewReq = &req{
		title:           title,
		fileIO:          file,
		content:         postText,
		nick:            name,
		authorIDSession: authorID.Value,
	}

	_, err = h.ctrl.NewPost(ctx, NewReq)

	if err != nil {
		h.HandleError(w, 400)
		return
	}

	log.Print("this new request:   ", NewReq)
}

func (r *req) GetTitle() string {
	return r.title
}

func (r *req) GetPostContent() string {
	return r.content
}

func (r *req) GetImage() io.Reader {
	return r.fileIO
}

func (r *req) GetName() string {
	return r.nick
}
func (r *req) GetAuthorIDSession() string {
	return r.authorIDSession
}
