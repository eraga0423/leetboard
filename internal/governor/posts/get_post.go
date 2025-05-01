package posts_governor

import (
	"context"
	"log"
	"time"

	"1337b0rd/internal/types/controller"
	"1337b0rd/internal/types/database"
)

type respOnePostGov struct {
	onePost  onePostGov
	comments []comment
}

func (r *respOnePostGov) GetOnePost() controller.ItemOnePost { return &r.onePost }
func (r *respOnePostGov) GetComments() []controller.Comment {
	result := make([]controller.Comment, len(r.comments))
	for i, c := range r.comments {
		result[i] = &c
	}
	return result
}

func (r *onePostGov) GetTitle() string                 { return r.title }
func (r *onePostGov) GetContent() string               { return r.content }
func (r *onePostGov) GetImageURL() string              { return r.imageURL }
func (r *onePostGov) GetPostTime() time.Time           { return r.postTime }
func (r *onePostGov) GetAuthorPost() controller.Author { return &r.authorPost }

func (r *authorGov) GetName() string      { return r.name }
func (r *authorGov) GetImageURL() string  { return r.imageURL }
func (r *authorGov) GetSessionID() string { return r.sessionID }

func (r *comment) GetParent() controller.OneComment { return &r.parentComment }
func (r *comment) GetChildren() []controller.OneComment {
	result := make([]controller.OneComment, len(r.childrenComment))
	for i, o := range r.childrenComment {
		result[i] = &o
	}
	return result
}

func (r *oneComment) GetCommentID() int            { return r.commentID }
func (r *oneComment) GetPostID() int               { return r.postID }
func (r *oneComment) GetAuthor() controller.Author { return &r.author }
func (r *oneComment) GetCommentContent() string    { return r.commentContent }
func (r *oneComment) GetCommentImage() string      { return r.commentImage }
func (r *oneComment) GetCommentTime() time.Time    { return r.commentTime }

type comment struct {
	parentComment   oneComment
	childrenComment []oneComment
}

type oneComment struct {
	commentID      int
	postID         int
	author         authorGov
	commentContent string
	commentImage   string
	commentTime    time.Time
}
type authorGov struct {
	name      string
	imageURL  string
	sessionID string
}
type onePostGov struct {
	title      string
	content    string
	imageURL   string
	postTime   time.Time
	authorPost authorGov
}

type newPostReq struct {
	postID int
}

func (p *PostsGovernor) OnePostGov(req controller.OnePostReq, ctx context.Context) (controller.OnePostResp, error) {
	request := newPostReq{
		postID: req.GetPostID(),
	}
	resp, err := p.db.OnePost(ctx, &request)
	if err != nil {
		log.Print("dir: postgres,  method: onePost, error:  ", err.Error())
		return nil, err
	}

	newRespPost := newResponsePost(resp)
	respComments := resp.GetComments()
	var newRespComment []comment
	for _, v := range respComments {
		newParentComment := newResponseParentComment(v.GetParent())
		newChildComments := newResponseChildComments(v.GetChildren())
		newRespComment = append(newRespComment, comment{
			parentComment:   newParentComment,
			childrenComment: newChildComments,
		})
	}

	newRespOnePostGov := respOnePostGov{
		onePost:  newRespPost,
		comments: newRespComment,
	}

	return &newRespOnePostGov, nil
}

func (n *newPostReq) ReqPostID() int {
	return n.postID
}

func newResponsePost(resp database.OnePostResp) onePostGov {
	respAuthPost := resp.GetOnePost().GetAuthorPost()
	respOnePost := resp.GetOnePost()
	newOnePost := onePostGov{
		title:    respOnePost.GetTitle(),
		content:  respOnePost.GetPostContent(),
		imageURL: respOnePost.GetPostUrlImage(),
		postTime: respOnePost.GetPostTime(),
		authorPost: authorGov{
			name:      respAuthPost.GetName(),
			imageURL:  respAuthPost.GetImageURL(),
			sessionID: respAuthPost.GetSessionID(),
		},
	}
	return newOnePost
}

func newResponseParentComment(parentComment database.OneComment) oneComment {
	respComment := oneComment{
		commentID: parentComment.GetCommentID(),
		postID:    parentComment.GetPostID(),
		author: authorGov{
			name:      parentComment.GetAuthor().GetName(),
			imageURL:  parentComment.GetAuthor().GetImageURL(),
			sessionID: parentComment.GetAuthor().GetSessionID(),
		},
		commentContent: parentComment.GetCommentContent(),
		commentTime:    parentComment.GetCommentTime(),
	}
	return respComment
}

func newResponseChildComments(childComments []database.OneComment) []oneComment {
	newComments := make([]oneComment, len(childComments))
	for _, comment := range childComments {
		newComments = append(newComments, oneComment{
			commentID: comment.GetCommentID(),
			postID:    comment.GetPostID(),
			author: authorGov{
				name:      comment.GetAuthor().GetName(),
				imageURL:  comment.GetAuthor().GetImageURL(),
				sessionID: comment.GetAuthor().GetSessionID(),
			},
			commentContent: comment.GetCommentContent(),
			commentTime:    comment.GetCommentTime(),
		})
	}
	return newComments
}
