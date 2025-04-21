package posts_handler

import (
	"1337b0rd/internal/constants"
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
	postID := r.FormValue("post_id") //////////////////////////////////////////////////////
	parentCommentID := r.FormValue("parent_id")
	sessionID, err := r.Cookie(constants.SessionIDKey)
	if err != nil {
		h.HandleError(w, http.StatusInternalServerError)
		return
	}
	avatarImageURl, err := r.Cookie(constants.ImageURL)
	if err != nil {
		h.HandleError(w, http.StatusInternalServerError)
		return
	}
	avatarName, err := r.Cookie(constants.Name)
	if err != nil {
		h.HandleError(w, http.StatusInternalServerError)
		return
	}
	commentContent := r.FormValue("comment_content")
	file, header, err := r.FormFile("image")
	if err != nil {
		h.HandleError(w, http.StatusInternalServerError)
		return
	}
	defer file.Close()
	newResp := respNewComment{
		avatarImageURL:  avatarImageURl.Value,
		postID:          postID,
		parentCommentID: parentCommentID,
		sessionID:       sessionID.Value,
		name:            avatarName.Value,
		content:         commentContent,

		image: metadataComment{
			fileIO:      file,
			objectSize:  header.Size,
			contentType: header.Header.Get("Content-Type"),
		}}
	_, err = h.ctrl.NewComment(&newResp)
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
