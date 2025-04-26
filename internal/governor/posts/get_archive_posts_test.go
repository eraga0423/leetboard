package posts_governor

import (
	"context"
	"testing"
	"time"

	"1337b0rd/internal/types/controller"
	"1337b0rd/internal/types/database"
)

type mockListPostsArchiveResp struct {
	posts []controller.ItemArchivePostsResp
}

func (m *mockListPostsArchiveResp) GetArchiveList() []controller.ItemArchivePostsResp {
	return m.posts
}

type mockArchivePost struct {
	postID       int
	title        string
	postContent  string
	postImageURL string
	postTime     time.Time
}

func (m *mockArchivePost) GetPostID() int          { return m.postID }
func (m *mockArchivePost) GetTitle() string        { return m.title }
func (m *mockArchivePost) GetPostContent() string  { return m.postContent }
func (m *mockArchivePost) GetPostImageURL() string { return m.postImageURL }
func (m *mockArchivePost) GetPostTime() time.Time  { return m.postTime }

type mockDBArchive struct {
	resp database.ListPostsArchiveResp
	err  error
}

func (m *mockDBArchive) ListArchivePosts() (controller.ListArchivePostsResp, error) {
	return m.resp, m.err
}

func TestListArchivePosts_Succes(t *testing.T) {
	ctx := context.Background()
	pg := &PostsGovernor{
		db: &mockDBArchive{},
	}
	resp, err := pg.ListArchivePosts(ctx)
}
