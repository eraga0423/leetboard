package posts_governor

import (
	"1337b0rd/internal/types/controller"
	"context"
	"time"
)

type listArchivePostsResp struct {
	archivePosts []onePostArchive
}
type onePostArchive struct {
	postID       int
	title        string
	postContent  string
	postImageURL string
	postTime     time.Time
}

func (g *PostsGovernor) ListArchivePosts(_ context.Context) (controller.ListArchivePostsResp, error) {
	resp, err := g.db.ListArchivePosts()
	if err != nil {
		return nil, err
	}
	newArchive := listArchivePostsResp{}
	list := resp.GetArchiveList()
	for _, archiveResp := range list {
		newArchive.archivePosts = append(newArchive.archivePosts, onePostArchive{
			postID:       archiveResp.GetPostID(),
			title:        archiveResp.GetTitle(),
			postContent:  archiveResp.GetPostContent(),
			postImageURL: archiveResp.GetPostImageURL(),
			postTime:     archiveResp.GetPostTime(),
		})
	}

	return &newArchive, nil
}

func (a *listArchivePostsResp) GetList() []controller.ItemArchivePostsResp {
	archives := make([]controller.ItemArchivePostsResp, len(a.archivePosts))
	for i, post := range a.archivePosts {
		archives[i] = &post

	}
	return archives
}

func (a *onePostArchive) GetPostID() int {
	return a.postID
}
func (a *onePostArchive) GetTitle() string {
	return a.title
}
func (a *onePostArchive) GetPostContent() string {
	return a.postContent
}
func (a *onePostArchive) GetPostImageURL() string {
	return a.postImageURL
}
func (a *onePostArchive) GetPostTime() time.Time {
	return a.postTime
}
