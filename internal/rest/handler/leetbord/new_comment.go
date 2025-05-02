package posts_handler

import (
	"fmt"
	"log/slog"
	"mime/multipart"
	"net/http"

	"1337b0rd/internal/types/controller"
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
	slog.Info("new comment", "", "")
	ctx := r.Context()
	postID := r.PathValue("id") //////////////////////////////////////////////////////
	slog.Info("postID", "val", postID)
	parentCommentID := r.FormValue("parent_id")
	slog.Info("parentID", "val", parentCommentID)
	commentContent := r.FormValue("comment_content")
	slog.Info("content", "val", commentContent)
	newCookie, err := parseCookie(r)
	if err != nil {
		slog.Any("parse cookie error: %w", err)
		h.HandleError(w, http.StatusInternalServerError)
		return
	}
	newImage, err := checkFile(r)
	if err != nil {
		slog.Any("check file error: %w", err)
		h.HandleError(w, http.StatusBadRequest)
		return
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
		slog.Any("new comment error: %w", err)
		h.HandleError(w, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/post/%s", postID), http.StatusSeeOther)
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
