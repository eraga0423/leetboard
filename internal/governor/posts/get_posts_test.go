package posts_governor

import (
	"context"
	"testing"
	"time"

	"1337b0rd/internal/types/database"
)

type mockListPostsResp struct {
	posts []database.ItemPostsResp
}

func (m *mockListPostsResp) GetList() []database.ItemPostsResp {
	return m.posts
}

type mockPost struct {
	postID       int
	title        string
	postContent  string
	postImageURL string
	postTime     time.Time
}

func (m *mockPost) GetPostID() int          { return m.postID }
func (m *mockPost) GetTitle() string        { return m.title }
func (m *mockPost) GetPostContent() string  { return m.postContent }
func (m *mockPost) GetPostImageURL() string { return m.postImageURL }
func (m *mockPost) GetPostTime() time.Time  { return m.postTime }

type mockDB struct {
	resp database.ListPostsResp
	err  error
}

func (m *mockDB) ListPosts() (database.ListPostsResp, error) {
	return m.resp, m.err
}

func TestPostsGovernor_List_Succes(t *testing.T) {
	mockPost := &mockPost{
		postID:       2,
		title:        "",
		postContent:  "",
		postImageURL: "",
		postTime:     time.Now(),
	}
	mockResp := &mockListPostsResp{posts: []database.ItemPostsResp{mockPost}}

	mockDB := &mockDB{resp: mockResp}
	gov := &PostsGovernor{db: mockDB}
	res, err := gov.ListPosts(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	list := res.GetList()
	if len(list) != 1 {
		t.Errorf("expected 1 post, got %d", len(list))
	}
	if list[0].GetPostID() != mockPost.GetPostID() {
		t.Errorf("expected post ID %d, got %d", mockPost.GetPostID(), list[0].GetPostID())
	}
}

func (m *mockDB) InserCartoonCharacters(database.InsertCharacters) error       { return nil }
func (m *mockDB) ListArchivePosts() (database.ListPostsArchiveResp, error)     { return nil, nil }
func (m *mockDB) ListCharacters() (database.ResponseCharacters, error)         { return nil, nil }
func (m *mockDB) CreatePost(database.NewPostReq) (database.NewPostResp, error) { return nil, nil }
func (m *mockDB) CreateComment(database.NewReqComment) error                   { return nil }
func (m *mockDB) OneArchivePost(database.ArchiveOnePostReq) (database.ArchiveOnePostResp, error) {
	return nil, nil
}
func (m *mockDB) OnePost(database.OnePostReq) (database.OnePostResp, error) { return nil, nil }
func (m *mockDB) UpdateCharacters(database.RequestCharacters) error         { return nil }
