package posts_handler

import (
	"log/slog"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"1337b0rd/internal/constants"
	"1337b0rd/internal/types/controller"
)

type archiveReqID struct {
	id int
}

type RespOnePostArchive struct {
	OnePost  OnePostArchiveId
	Comments []CommentArchive
}
type CommentArchive struct {
	ParentComment   OneCommentArchive
	ChildrenComment []OneCommentArchive
}

type OneCommentArchive struct {
	CommentID      int
	PostID         int
	Author         AuthorArchive
	CommentContent string
	CommentImage   string
	CommentTime    time.Time
}
type AuthorArchive struct {
	Name      string
	ImageURL  string
	SessionID string
}
type OnePostArchiveId struct {
	Title      string
	Content    string
	ImageURL   string
	PostTime   time.Time
	AuthorPost AuthorArchive
}

func (h *PostsHandler) GetPostIDArchive(w http.ResponseWriter, r *http.Request) {
	slog.Info("this metod get post id archive", "dir", "rest")
	ctx := r.Context()
	tmpl := template.Must(template.ParseFiles(constants.ArchivePost))
	id := r.PathValue("id")
	slog.Info(id)
	intID, err := strconv.Atoi(id)
	if err != nil {
		h.HandleError(w, http.StatusInternalServerError)
		return
	}
	var req archiveReqID
	req = archiveReqID{
		id: intID,
	}

	res, err := h.ctrl.OneArchivePostGov(&req, ctx)
	if err != nil {
		h.HandleError(w, http.StatusInternalServerError)
		slog.Info("method", "gov", err)
		return
	}
	respMain := RespOnePostArchive{}
	comments := res.GetComments()
	postResp := newArchivePost(res)
	respMain.OnePost = postResp
	var newRespComment []CommentArchive
	for _, v := range comments {
		parentResp := newArchiveParentComment(v.GetParent())
		childResp := newArchiveChildrenComment(v.GetChildren())
		newRespComment = append(newRespComment, CommentArchive{
			ParentComment:   parentResp,
			ChildrenComment: childResp,
		})
	}
	respMain.Comments = newRespComment
	err = tmpl.Execute(w, respMain)
	if err != nil {
		h.HandleError(w, http.StatusInternalServerError)
		slog.Info("method", "front", err)
		return
	}
}

func (a *archiveReqID) GetPostID() int {
	return a.id
}

func newArchivePost(req controller.ArchiveOnePostResp) OnePostArchiveId {
	respAuthPost := req.GetOnePost().GetAuthorPost()
	respOnePost := req.GetOnePost()
	newOnePost := OnePostArchiveId{
		Title:    respOnePost.GetTitle(),
		Content:  respOnePost.GetPostContent(),
		ImageURL: respAuthPost.GetImageURL(),
		PostTime: respOnePost.GetPostTime(),
		AuthorPost: AuthorArchive{
			Name:      respAuthPost.GetName(),
			ImageURL:  respAuthPost.GetImageURL(),
			SessionID: respAuthPost.GetSessionID(),
		},
	}
	return newOnePost
}

func newArchiveParentComment(parent controller.ArchiveOneComment) OneCommentArchive {
	respComment := OneCommentArchive{
		CommentID: parent.GetCommentID(),
		PostID:    parent.GetPostID(),
		Author: AuthorArchive{
			Name:      parent.GetAuthor().GetName(),
			ImageURL:  parent.GetAuthor().GetImageURL(),
			SessionID: parent.GetAuthor().GetSessionID(),
		},
		CommentContent: parent.GetCommentContent(),
		CommentTime:    parent.GetCommentTime(),
	}
	return respComment
}

func newArchiveChildrenComment(child []controller.ArchiveOneComment) []OneCommentArchive {
	newComment := make([]OneCommentArchive, 0, len(child))
	for _, comment := range child {
		newComment = append(newComment, OneCommentArchive{
			CommentID: comment.GetCommentID(),
			PostID:    comment.GetPostID(),
			Author: AuthorArchive{
				Name:      comment.GetAuthor().GetName(),
				ImageURL:  comment.GetAuthor().GetImageURL(),
				SessionID: comment.GetAuthor().GetSessionID(),
			},
			CommentContent: comment.GetCommentContent(),
			CommentTime:    comment.GetCommentTime(),
		},
		)
	}
	return newComment
}
