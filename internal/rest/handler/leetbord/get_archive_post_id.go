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

type respOnePostArchive struct {
	onePost  onePostArchive
	comments []commentArchive
}
type commentArchive struct {
	parentComment   oneCommentArchive
	childrenComment []oneCommentArchive
}

type oneCommentArchive struct {
	commentID      int
	postID         int
	author         authorArchive
	commentContent string
	commentImage   string
	commentTime    time.Time
}
type authorArchive struct {
	name      string
	imageURL  string
	sessionID string
}
type onePostArchive struct {
	title      string
	content    string
	imageURL   string
	postTime   time.Time
	authorPost authorArchive
}

func (h *PostsHandler) GetPostIDArchive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tmpl := template.Must(template.ParseFiles(constants.Post))
	id := r.PathValue("id")
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
	comments := res.GetComments()
	postResp := newArchivePost(res)
	var newRespComment []commentArchive
	for _, v := range comments {
		parentResp := newArchiveParentComment(v.GetParent())
		childResp := newArchiveChildrenComment(v.GetChildren())
		newRespComment = append(newRespComment, commentArchive{
			parentComment:   parentResp,
			childrenComment: childResp,
		})
		respMain := respOnePostArchive{
			onePost:  postResp,
			comments: newRespComment,
		}
		err = tmpl.Execute(w, respMain)
		if err != nil {
			h.HandleError(w, http.StatusInternalServerError)
			slog.Info("method", "front", err)
			return
		}
	}
}

func (a *archiveReqID) GetPostID() int {
	return a.id
}

func newArchivePost(req controller.ArchiveOnePostResp) onePostArchive {
	respAuthPost := req.GetOnePost().GetAuthorPost()
	respOnePost := req.GetOnePost()
	newOnePost := onePostArchive{
		title:    respOnePost.GetTitle(),
		content:  respOnePost.GetPostContent(),
		imageURL: respAuthPost.GetImageURL(),
		postTime: respOnePost.GetPostTime(),
		authorPost: authorArchive{
			name:      respAuthPost.GetName(),
			imageURL:  respAuthPost.GetImageURL(),
			sessionID: respAuthPost.GetSessionID(),
		},
	}
	return newOnePost
}

func newArchiveParentComment(parent controller.ArchiveOneComment) oneCommentArchive {
	respComment := oneCommentArchive{
		commentID: parent.GetCommentID(),
		postID:    parent.GetPostID(),
		author: authorArchive{
			name:      parent.GetAuthor().GetName(),
			imageURL:  parent.GetAuthor().GetImageURL(),
			sessionID: parent.GetAuthor().GetSessionID(),
		},
		commentContent: parent.GetCommentContent(),
		commentTime:    parent.GetCommentTime(),
	}
	return respComment
}

func newArchiveChildrenComment(child []controller.ArchiveOneComment) []oneCommentArchive {
	newComment := make([]oneCommentArchive, len(child))
	for _, comment := range child {
		newComment = append(newComment, oneCommentArchive{
			commentID: comment.GetCommentID(),
			postID:    comment.GetPostID(),
			author: authorArchive{
				name:      comment.GetAuthor().GetName(),
				imageURL:  comment.GetAuthor().GetImageURL(),
				sessionID: comment.GetAuthor().GetSessionID(),
			},
			commentContent: comment.GetCommentContent(),
			commentTime:    comment.GetCommentTime(),
		},
		)
	}
	return newComment
}
