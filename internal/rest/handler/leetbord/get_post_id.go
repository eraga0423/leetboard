package posts_handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"1337b0rd/internal/constants"
	"1337b0rd/internal/types/controller"
)

type reqID struct {
	id int
}

type RespOnePost struct {
	OnePost  OnePost
	Comments []Comment
}
type Comment struct {
	ParentComment   OneComment
	ChildrenComment []OneComment
}

type OneComment struct {
	CommentID      int
	PostID         int
	Author         Author
	CommentContent string
	CommentImage   string
	CommentTime    time.Time
}
type Author struct {
	Name      string
	ImageURL  string
	SessionID string
}
type OnePost struct {
	PostID     int
	Title      string
	Content    string
	ImageURL   string
	PostTime   time.Time
	AuthorPost Author
}

func (h *PostsHandler) GetPostID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This GET /post/id")
	ctx := r.Context()
	tmpl := template.Must(template.ParseFiles(constants.Post))
	id := r.PathValue("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		slog.Info("method", "atoi", err)
		h.HandleError(w, http.StatusInternalServerError)
		return
	}
	req := new(reqID)
	req = &reqID{id: intID}
	data, err := h.ctrl.OnePostGov(req, ctx)
	if err != nil {
		slog.Error("method", "gov", err)
		h.HandleError(w, http.StatusInternalServerError)
		return
	}
	var newComment []Comment
	newPost := newPost(data, intID)
	for _, v := range data.GetComments() {
		newParent := newParentComment(v.GetParent())
		newChild := newChildComments(v.GetChildren())
		newComment = append(newComment, Comment{
			ChildrenComment: newChild,
			ParentComment:   newParent,
		})
	}
	mainResp := RespOnePost{
		OnePost:  newPost,
		Comments: newComment,
	}

	err = tmpl.Execute(w, mainResp)
	if err != nil {
		h.HandleError(w, http.StatusInternalServerError)
		slog.Info("method", "front", err)
		return
	}
}

func (i *reqID) GetPostID() int {
	return i.id
}

func newPost(req controller.OnePostResp, postID int) OnePost {
	respAuthPost := req.GetOnePost().GetAuthorPost()
	respOnePost := req.GetOnePost()
	newOnePost := OnePost{
		Title:    respOnePost.GetTitle(),
		Content:  respOnePost.GetContent(),
		ImageURL: respOnePost.GetImageURL(),
		PostTime: respOnePost.GetPostTime(),
		AuthorPost: Author{
			Name:      respAuthPost.GetName(),
			ImageURL:  respAuthPost.GetImageURL(),
			SessionID: respAuthPost.GetSessionID(),
		},
		PostID: postID,
	}
	return newOnePost
}

func newParentComment(parent controller.OneComment) OneComment {
	respComment := OneComment{
		CommentID: parent.GetCommentID(),
		PostID:    parent.GetPostID(),
		Author: Author{
			Name:      parent.GetAuthor().GetName(),
			ImageURL:  parent.GetAuthor().GetImageURL(),
			SessionID: parent.GetAuthor().GetSessionID(),
		},
		CommentContent: parent.GetCommentContent(),
		CommentTime:    parent.GetCommentTime(),
	}
	return respComment
}

func newChildComments(child []controller.OneComment) []OneComment {
	newComment := make([]OneComment, 0, len(child))
	for _, com := range child {
		newComment = append(newComment, OneComment{
			CommentID: com.GetCommentID(),
			PostID:    com.GetPostID(),
			Author: Author{
				Name:      com.GetAuthor().GetName(),
				ImageURL:  com.GetAuthor().GetImageURL(),
				SessionID: com.GetAuthor().GetSessionID(),
			},
			CommentContent: com.GetCommentContent(),
			CommentTime:    com.GetCommentTime(),
		})
	}
	return newComment
}
