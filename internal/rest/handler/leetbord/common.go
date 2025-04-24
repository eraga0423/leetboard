package posts_handler

import (
	"1337b0rd/internal/constants"
	"1337b0rd/internal/types/controller"
	"errors"
	"mime/multipart"
	"net/http"
)

type PostsHandler struct {
	ctrl controller.Controller
	// logger *log.Logger
}

func New(ctrl controller.Controller) *PostsHandler {
	return &PostsHandler{ctrl: ctrl}
}

type metaData struct {
	fileIO      multipart.File
	objectSize  int64
	contentType string
}

func checkFile(r *http.Request) (metaData, error) {

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return metaData{}, err
	}
	file, header, err := r.FormFile("image")

	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			return metaData{
				fileIO:      nil,
				objectSize:  0,
				contentType: "",
			}, nil
		} else {
			return metaData{}, err
		}

	}
	if file == nil {
		defer file.Close()
	}
	contentType := header.Header.Get("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/png" {
		return metaData{}, errors.New("content type not supported")
	}
	return metaData{
		fileIO:      file,
		objectSize:  header.Size,
		contentType: contentType,
	}, nil
}

type respNewTemp struct {
	avatarImageURL string
	sessionID      string
	name           string
}

func parseCookie(r *http.Request) (respNewTemp, error) {
	sessionID, err := r.Cookie(constants.SessionIDKey)
	if err != nil {
		return respNewTemp{}, err
	}
	avatarImageURl, err := r.Cookie(constants.ImageURL)
	if err != nil {
		return respNewTemp{}, err
	}
	avatarName, err := r.Cookie(constants.Name)
	if err != nil {
		return respNewTemp{}, err
	}
	return respNewTemp{
		sessionID:      sessionID.Value,
		avatarImageURL: avatarImageURl.Value,
		name:           avatarName.Value,
	}, nil
}
