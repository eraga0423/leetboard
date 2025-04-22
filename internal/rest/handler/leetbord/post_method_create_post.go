package posts_handler

import (
	"log"
	"mime/multipart"
	"net/http"

	"1337b0rd/internal/types/controller"
)

type req struct {
	authorIDSession string
	formName        string
	defaultName     string
	avatarImageURL  string
	title           string
	content         string
	fileData        metadata
}
type metadata struct {
	fileIO      multipart.File
	objectSize  int64
	contentType string
}

func (h *PostsHandler) PostMethodCreatePost(w http.ResponseWriter, r *http.Request) {
	log.Print("This post /create")
	ctx := r.Context()

	name := r.FormValue("name")
	title := r.FormValue("title")
	postText := r.FormValue("post")

	newFile, err := checkFile(r)
	if err != nil {
		h.HandleError(w, http.StatusBadRequest)
		return
	}
	newCookie, err := parseCookie(r)
	if err != nil {
		h.HandleError(w, http.StatusInternalServerError)
		return
	}

	NewReq := &req{
		title:           title,
		content:         postText,
		formName:        name,
		defaultName:     newCookie.name,
		avatarImageURL:  newCookie.avatarImageURL,
		authorIDSession: newCookie.sessionID,
		fileData: metadata{
			fileIO:      newFile.fileIO,
			objectSize:  newFile.objectSize,
			contentType: newFile.contentType,
		},
	}

	_, err = h.ctrl.NewPost(ctx, NewReq)
	if err != nil {
		h.HandleError(w, 400)
		return
	}

	log.Print("this new request:   ", NewReq)
}

func (r *req) GetTitle() string               { return r.title }
func (r *req) GetPostContent() string         { return r.content }
func (r *req) GetFormName() string            { return r.formName }
func (r *req) GetAuthorIDSession() string     { return r.authorIDSession }
func (r *req) GetDefaultName() string         { return r.defaultName }
func (r *req) GetAvatarImageURL() string      { return r.avatarImageURL }
func (m *metadata) GetFileIO() multipart.File { return m.fileIO }
func (m *metadata) GetObjectSize() int64      { return m.objectSize }
func (m *metadata) GetContentType() string    { return m.contentType }

func (r *req) GetImage() controller.ItemMetaData {
	return &r.fileData
}
