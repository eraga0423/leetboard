package posts_handler

import (
	"fmt"
	"io"
	"net/http"
)

type Req struct {
	Title    string `json:"title"`
	Post     string `json:"post"`
	Nick     string `json:"name"`
	FileByte []byte
}

func (h *PostsHandler) PostMethodCreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This post /create")
	var NewReq Req
	name := r.FormValue("name")
	title := r.FormValue("title")
	postText := r.FormValue("post")
	file, _, err := r.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		http.Error(w, "Не удалось добавить файл", http.StatusBadRequest)
		return
	}
	var newFile []byte
	if file != nil {
		newFile, err = io.ReadAll(file)
		if err != nil {
			return
		}
		defer file.Close()
	}
	NewReq = Req{
		Title:    title,
		FileByte: newFile,
		Post:     postText,
		Nick:     name,
	}
	// здесь будет отправка этих данных в governor

	fmt.Println("asdasd", NewReq)
}
