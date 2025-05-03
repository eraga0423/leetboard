package posts_governor

import (
	"1337b0rd/internal/types/controller"
	"testing"
	"time"
)

// Моки для OnePostReq
type MockOnePostReq struct {
	PostID int
}

func (m *MockOnePostReq) GetPostID() int {
	return m.PostID
}

// Моки для Author
type MockAuthor struct {
	Name      string
	ImageURL  string
	SessionID string
}

func (m *MockAuthor) GetName() string {
	return m.Name
}

func (m *MockAuthor) GetImageURL() string {
	return m.ImageURL
}

func (m *MockAuthor) GetSessionID() string {
	return m.SessionID
}

// Моки для ItemOnePost
type MockItemOnePost struct {
	Title      string
	Content    string
	ImageURL   string
	PostTime   time.Time
	AuthorPost controller.Author
}

func (m *MockItemOnePost) GetTitle() string {
	return m.Title
}

func (m *MockItemOnePost) GetContent() string {
	return m.Content
}

func (m *MockItemOnePost) GetImageURL() string {
	return m.ImageURL
}

func (m *MockItemOnePost) GetPostTime() time.Time {
	return m.PostTime
}

func (m *MockItemOnePost) GetAuthorPost() controller.Author {
	return m.AuthorPost
}

// Моки для OneComment
type MockOneComment struct {
	CommentID      int
	PostID         int
	Author         controller.Author
	CommentContent string
	CommentImage   string
	CommentTime    time.Time
}

func (m *MockOneComment) GetCommentID() int {
	return m.CommentID
}

func (m *MockOneComment) GetPostID() int {
	return m.PostID
}

func (m *MockOneComment) GetAuthor() controller.Author {
	return m.Author
}

func (m *MockOneComment) GetCommentContent() string {
	return m.CommentContent
}

func (m *MockOneComment) GetCommentImage() string {
	return m.CommentImage
}

func (m *MockOneComment) GetCommentTime() time.Time {
	return m.CommentTime
}

// Моки для Comment
type MockComment struct {
	ParentComment   controller.OneComment
	ChildrenComment []controller.OneComment
}

func (m *MockComment) GetParent() controller.OneComment {
	return m.ParentComment
}

func (m *MockComment) GetChildren() []controller.OneComment {
	return m.ChildrenComment
}

// Моки для OnePostResp
type MockOnePostResp struct {
	OnePost  controller.ItemOnePost
	Comments []controller.Comment
}

func (m *MockOnePostResp) GetOnePost() controller.ItemOnePost {
	return m.OnePost
}

func (m *MockOnePostResp) GetComments() []controller.Comment {
	return m.Comments
}

func TestMocks(t *testing.T) {
	// Создание моков для теста
	mockReq := &MockOnePostReq{PostID: 123}
	mockAuthor := &MockAuthor{Name: "Author Name", ImageURL: "http://example.com/image.jpg", SessionID: "abc123"}
	mockItemPost := &MockItemOnePost{
		Title:      "Mock Post Title",
		Content:    "This is a mock post",
		ImageURL:   "http://example.com/post-image.jpg",
		PostTime:   time.Now(),
		AuthorPost: mockAuthor,
	}
	mockComment := &MockOneComment{
		CommentID:      1,
		PostID:         123,
		Author:         mockAuthor,
		CommentContent: "This is a comment",
		CommentImage:   "http://example.com/comment-image.jpg",
		CommentTime:    time.Now(),
	}
	mockComments := []controller.Comment{&MockComment{ParentComment: mockComment}}

	// Создание мок-ответа
	mockResp := &MockOnePostResp{
		OnePost:  mockItemPost,
		Comments: mockComments,
	}

	// Тестирование получения поста
	if mockResp.GetOnePost().GetTitle() != "Mock Post Title" {
		t.Errorf("Expected 'Mock Post Title', got %s", mockResp.GetOnePost().GetTitle())
	}

	// Тестирование комментариев
	if len(mockResp.GetComments()) != 1 {
		t.Errorf("Expected 1 comment, got %d", len(mockResp.GetComments()))
	}

	// Тестирование авторов комментариев
	if mockResp.GetComments()[0].GetParent().GetAuthor().GetName() != "Author Name" {
		t.Errorf("Expected 'Author Name', got %s", mockResp.GetComments()[0].GetParent().GetAuthor().GetName())
	}

	// Тестирование PostID
	if mockReq.GetPostID() != 123 {
		t.Errorf("Expected PostID 123, got %d", mockReq.GetPostID())
	}
}
