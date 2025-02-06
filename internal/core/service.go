package core

type PostService struct {
	repo PostRepository
}

func NewPostService(repo PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(content, image string) (*Post, error) {
	post := &Post{Content: content, Image: image}
	return s.repo.Save(post)
}
