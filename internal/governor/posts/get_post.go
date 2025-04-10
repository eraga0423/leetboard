package posts_governor

import (
	"1337b0rd/internal/types/controller"
	"1337b0rd/internal/types/database"
	"context"
	"log"
	"time"
)

type respOnePostGov struct {
	onePost  onePostGov
	comments []comment
}

//func (r respOnePostGov) GetOnePost() controller.ItemOnePost {
//	var itemOnePost controller.ItemOnePost
//	itemOnePost
//}

//
//func (r respOnePostGov) GetComments() []controller.Comment {
//	//TODO implement me
//	panic("implement me")
//}

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

func (p PostsGovernor) OnePostGov(req controller.OnePostReq, ctx context.Context) (controller.OnePostResp, error) {

	request := newPostReq{
		postID: req.GetPostID(),
	}
	resp, err := p.db.OnePost(request)
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

	return newRespOnePostGov, nil

}

func (n newPostReq) ReqPostID() int {
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
