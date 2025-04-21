package posts_handler

import (
	"1337b0rd/internal/types/controller"
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

	newResp := &respNewComment{
		avatarImageURL:  newCookie.avatarImageURL,
		postID:          postID,
		parentCommentID: parentCommentID,
		sessionID:       newCookie.sessionID,
		name:            newCookie.name,
		content:         commentContent,
		image: metadataComment{
			fileIO:      newImage.fileIO,
			objectSize:  newImage.objectSize,
			contentType: newImage.contentType,
		},
	}
	_, err = h.ctrl.NewComment(newResp, ctx)
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
