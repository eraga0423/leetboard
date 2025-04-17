package posts_handler

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	"1337b0rd/internal/constants"
	"1337b0rd/internal/types/controller"
)

type req struct {
	authorIDSession string
	title           string
	content         string
	nick            string
	fileData        metadata
}
type metadata struct {
	fileIO      multipart.File
	objectSize  int64
	contentType string
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
	file, fileHeader, err := r.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		h.HandleError(w, http.StatusBadRequest)
		return
	}

	authorID, err := r.Cookie(constants.SessionIDKey)
	if err != nil {
		return
	}

	defer file.Close()

	NewReq = &req{
		title:           title,
		content:         postText,
		nick:            name,
		authorIDSession: authorID.Value,
		fileData: metadata{
			fileIO:      file,
			objectSize:  fileHeader.Size,
			contentType: fileHeader.Header.Get("Content-Type"),
		},
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

func (r *req) GetImage() controller.ItemMetaData {
	var result controller.ItemMetaData
	result = &r.fileData
	return result
}

func (r *req) GetName() string {
	return r.nick
}

func (r *req) GetAuthorIDSession() string {
	return r.authorIDSession
}

func (m *metadata) GetFileIO() multipart.File {
	return m.fileIO
}

func (m *metadata) GetObjectSize() int64 {
	return m.objectSize
}

func (m *metadata) GetContentType() string {
	return m.contentType
}
