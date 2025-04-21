package posts_handler

import (
	"1337b0rd/internal/constants"
	"1337b0rd/internal/types/controller"
	"errors"
	"mime/multipart"
	"net/http"
)

type metadataComment struct {
	fileIO      multipart.File
	objectSize  int64
	contentType string
}
type respNewComment struct {
	avatarImageURL  string
	postID          string
	parentCommentID string
	sessionID       string
	name            string
	content         string
	image           metadataComment
}

func (h *PostsHandler) NewComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postID := r.FormValue("post_id") //////////////////////////////////////////////////////
	parentCommentID := r.FormValue("parent_id")
	commentContent := r.FormValue("comment_content")
	newCookie, err := parseCookie(r)
	if err != nil {
		h.HandleError(w, http.StatusInternalServerError)
		return
	}
	newImage, err := checkFile(r)
	if err != nil {
		h.HandleError(w, http.StatusBadRequest)
	}

	newResp := respNewComment{
		avatarImageURL:  newCookie.avatarImageURL,
		postID:          postID,
		parentCommentID: parentCommentID,
		sessionID:       newCookie.sessionID,
		name:            newCookie.name,
		content:         commentContent,
		image:           newImage,
	}
	_, err = h.ctrl.NewComment(&newResp, ctx)
	if err != nil {
		h.HandleError(w, http.StatusInternalServerError)
		return
	}
}
func (c *respNewComment) GetPostID() string                           { return c.postID }
func (c *respNewComment) GetAvatarImageURL() string                   { return c.avatarImageURL }
func (c *respNewComment) GetAvatarName() string                       { return c.name }
func (c *respNewComment) GetParentCommentID() string                  { return c.parentCommentID }
func (c *respNewComment) GetSessionID() string                        { return c.sessionID }
func (c *respNewComment) GetContent() string                          { return c.content }
func (c *respNewComment) GetImageComment() controller.MetaDataComment { return &c.image }
func (m *metadataComment) GetFileIO() multipart.File                  { return m.fileIO }
func (m *metadataComment) GetObjectSize() int64                       { return m.objectSize }
func (m *metadataComment) GetContentType() string                     { return m.contentType }

func checkFile(r *http.Request) (metadataComment, error) {
	file, header, err := r.FormFile("image")
	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			return metadataComment{
				objectSize: 0,
			}, nil
		} else {
			return metadataComment{}, err
		}
		defer file.Close()
	}

	return metadataComment{
		fileIO:      file,
		objectSize:  header.Size,
		contentType: header.Header.Get("Content-Type"),
	}, nil
}

func parseCookie(r *http.Request) (respNewComment, error) {
	sessionID, err := r.Cookie(constants.SessionIDKey)
	if err != nil {
		return respNewComment{}, err
	}
	avatarImageURl, err := r.Cookie(constants.ImageURL)
	if err != nil {
		return respNewComment{}, err
	}
	avatarName, err := r.Cookie(constants.Name)
	if err != nil {
		return respNewComment{}, err
	}
	return respNewComment{
		sessionID:      sessionID.Value,
		avatarImageURL: avatarImageURl.Value,
		name:           avatarName.Value,
	}, nil
}
