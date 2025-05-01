package posts_governor

import (
	"context"
	"errors"
	"testing"
	"time"

	"1337b0rd/internal/types/database"
)

type mockListPostsArchiveResp struct {
	posts []database.ItemPostsArchiveResp
}

func (m *mockListPostsArchiveResp) GetArchiveList() []database.ItemPostsArchiveResp {
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

func (m *mockDBArchive) ListArchivePosts(context.Context) (database.ListPostsArchiveResp, error) {
	return m.resp, m.err
}

func TestPostGovernor_ListArchivePosts_Succes(t *testing.T) {
	mockPost := &mockArchivePost{
		postID:       1,
		title:        "title",
		postContent:  "content",
		postImageURL: "http??sdsfsf",
		postTime:     time.Now(),
	}
	mockresp := &mockListPostsArchiveResp{
		posts: []database.ItemPostsArchiveResp{mockPost},
	}
	mockDB := &mockDBArchive{resp: mockresp}
	gov := &PostsGovernor{db: mockDB}
	res, err := gov.ListArchivePosts(context.Background())
	if err != nil {
		t.Fatalf("expected no errror, got %v", err)
	}
	list := res.GetList()
	if len(list) != 1 {
		t.Fatalf("expected 1 post, got %d", len(list))
	}
	if list[0].GetPostID() != mockPost.GetPostID() {
		t.Errorf("expected post ID %d, got %d", mockPost.GetPostID(), list[0].GetPostID())
	}
}

func TestPostsGovernor_ListArchivePosts_Error(t *testing.T) {
	expectedErr := errors.New("db error")
	mockDB := &mockDBArchive{err: expectedErr}
	gov := &PostsGovernor{db: mockDB}
	_, err := gov.ListArchivePosts(context.Background())
	if err == nil {
		t.Fatalf("expected an error, got none")
	}
	if err != expectedErr {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
}

func (m *mockDBArchive) CreateComment(context.Context, database.NewReqComment) (database.NewRespComment, error) {
	return nil, nil
}

func (m *mockDBArchive) ListPosts(context.Context) (database.ListPostsResp, error) {
	return nil, nil
}

func (m *mockDBArchive) CreatePost(context.Context, database.NewPostReq) (database.NewPostResp, error) {
	return nil, nil
}

func (m *mockDBArchive) OnePost(context.Context, database.OnePostReq) (database.OnePostResp, error) {
	return nil, nil
}

func (m *mockDBArchive) OneArchivePost(context.Context, database.ArchiveOnePostReq) (database.ArchiveOnePostResp, error) {
	return nil, nil
}

func (m *mockDBArchive) ListCharacters(context.Context) (database.ResponseCharacters, error) {
	return nil, nil
}

func (m *mockDBArchive) UpdateCharacters(context.Context, database.RequestCharacters) error {
	return nil
}

func (m *mockDBArchive) InserCartoonCharacters(context.Context, database.InsertCharacters) error {
	return nil
}
