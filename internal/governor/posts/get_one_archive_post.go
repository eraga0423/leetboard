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

func (g *PostsGovernor) OneArchivePostGov(onePostResp controller.ArchiveOnePostReq, ctx context.Context) (controller.ArchiveOnePostResp, error) {
	postID := archiveOnePostReq{}
	postID.postArchiveID = onePostResp.GetPostID()
	resp, err := g.db.OneArchivePost(&postID)
	if err != nil {
		log.Print("dir: postgres,  method: oneArchivePost, error:  ", err.Error())
		return nil, err
	}

	return nil, nil
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

func (a *archiveOnePostReq) ReqPostID() int {
	return a.postArchiveID
}
