package posts_governor

import (
	"1337b0rd/internal/types/controller"
	"1337b0rd/internal/types/database"
	"context"
	"log"
	"time"
)

type archiveOnePostResp struct {
	onePost  archiveRespOnePost
	comments []archiveComment
}
type archiveRespOnePost struct {
	title        string
	postContent  string
	postURLImage string
	postTime     time.Time
	authorPost   archiveRespOnePostAuthor
}
type archiveRespOnePostAuthor struct {
	name      string
	imageURL  string
	sessionID string
}
type archiveComment struct {
	parent   archiveOneComment
	children []archiveOneComment
}

type archiveOneComment struct {
	commentID       int
	postID          int
	author          archiveRespCommentAuthor
	commentContent  string
	commentImageURL string
	commentTime     time.Time
}
type archiveRespCommentAuthor struct {
	name      string
	imageURL  string
	sessionID string
}

type archiveOnePostReq struct {
	postArchiveID int
}

func (g *PostsGovernor) OneArchivePostGov(onePostReq controller.ArchiveOnePostReq, ctx context.Context) (controller.ArchiveOnePostResp, error) {
	postID := archiveOnePostReq{}
	postID.postArchiveID = onePostReq.GetPostID()
	resp, err := g.db.OneArchivePost(&postID)
	if err != nil {
		log.Print("dir: postgres,  method: oneArchivePost, error:  ", err.Error())
		return nil, err
	}
	newRespPostArchive := newOnePost(resp)
	respComment := resp.GetComments()
	var newRespComments []archiveComment
	for _, v := range respComment {
		newRespComments = append(newRespComments, archiveComment{
			parent:   newRespParentComment(v.GetParent()),
			children: newRespChildComment(v.GetChildren()),
		})
	}
	newRespPostArchiveGov := archiveOnePostResp{
		onePost:  newRespPostArchive,
		comments: newRespComments,
	}
	return &newRespPostArchiveGov, nil
}

func newOnePost(resp database.ArchiveOnePostResp) archiveRespOnePost {
	respAuthPost := resp.GetOnePost().GetAuthorPost()
	respOnePost := resp.GetOnePost()
	newOneArchivePost := archiveRespOnePost{
		title:        respOnePost.GetTitle(),
		postContent:  respOnePost.GetPostContent(),
		postURLImage: respOnePost.GetPostUrlImage(),
		postTime:     respOnePost.GetPostTime(),
		authorPost: archiveRespOnePostAuthor{
			name:      respAuthPost.GetName(),
			imageURL:  respAuthPost.GetImageURL(),
			sessionID: respAuthPost.GetSessionID(),
		},
	}
	return newOneArchivePost
}
func newRespParentComment(parentComment database.ArchiveOneComment) archiveOneComment {
	respComment := archiveOneComment{
		commentID: parentComment.GetCommentID(),
		postID:    parentComment.GetPostID(),
		author: archiveRespCommentAuthor{
			name:      parentComment.GetAuthor().GetName(),
			imageURL:  parentComment.GetAuthor().GetImageURL(),
			sessionID: parentComment.GetAuthor().GetSessionID(),
		},
		commentContent:  parentComment.GetCommentContent(),
		commentImageURL: parentComment.GetCommentImage(),
		commentTime:     parentComment.GetCommentTime(),
	}
	return respComment
}
func newRespChildComment(childComment []database.ArchiveOneComment) []archiveOneComment {
	var newArchiveChildComment []archiveOneComment
	for _, v := range childComment {
		newArchiveChildComment = append(newArchiveChildComment, archiveOneComment{
			commentID: v.GetCommentID(),
			postID:    v.GetPostID(),
			author: archiveRespCommentAuthor{
				name:      v.GetAuthor().GetName(),
				imageURL:  v.GetAuthor().GetImageURL(),
				sessionID: v.GetAuthor().GetSessionID(),
			},
			commentContent:  v.GetCommentContent(),
			commentImageURL: v.GetCommentImage(),
			commentTime:     v.GetCommentTime(),
		})

	}
	return newArchiveChildComment
}
func (a *archiveOnePostReq) ReqPostID() int { return a.postArchiveID }

func (a *archiveOnePostResp) GetOnePost() controller.ArchiveRespOnePost { return &a.onePost }

func (a *archiveOnePostResp) GetComments() []controller.ArchiveComment {
	res := make([]controller.ArchiveComment, len(a.comments))
	for i, v := range a.comments {
		res[i] = &v
	}
	return res
}

func (a *archiveRespOnePost) GetTitle() string        { return a.title }
func (a *archiveRespOnePost) GetPostContent() string  { return a.postContent }
func (a *archiveRespOnePost) GetPostUrlImage() string { return a.postURLImage }
func (a *archiveRespOnePost) GetPostTime() time.Time  { return a.postTime }
func (a *archiveRespOnePost) GetAuthorPost() controller.ArchiveRespOnePostAuthor {
	return &a.authorPost
}

func (c *archiveRespOnePostAuthor) GetName() string      { return c.name }
func (c *archiveRespOnePostAuthor) GetImageURL() string  { return c.imageURL }
func (c *archiveRespOnePostAuthor) GetSessionID() string { return c.sessionID }

func (a *archiveComment) GetParent() controller.ArchiveOneComment { return &a.parent }
func (a *archiveComment) GetChildren() []controller.ArchiveOneComment {
	res := make([]controller.ArchiveOneComment, len(a.children))
	for i, v := range a.children {
		res[i] = &v

	}

	return res
}

func (a *archiveOneComment) GetCommentID() int                              { return a.commentID }
func (a *archiveOneComment) GetPostID() int                                 { return a.postID }
func (a *archiveOneComment) GetAuthor() controller.ArchiveRespCommentAuthor { return &a.author }
func (a *archiveOneComment) GetCommentContent() string                      { return a.commentContent }
func (a *archiveOneComment) GetCommentImage() string                        { return a.commentImageURL }
func (a *archiveOneComment) GetCommentTime() time.Time                      { return a.commentTime }

func (a *archiveRespCommentAuthor) GetName() string      { return a.name }
func (a *archiveRespCommentAuthor) GetImageURL() string  { return a.imageURL }
func (a *archiveRespCommentAuthor) GetSessionID() string { return a.imageURL }
