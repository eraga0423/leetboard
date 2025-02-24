package governor

import "1337b0rd/internal/governor/posts"

type Governor struct {
	*posts.Posts
}

func New() *Governor {
	return &Governor{
		Posts: new(posts.Posts),
	}
}
