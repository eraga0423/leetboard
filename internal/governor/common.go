package governor

import posts_governor "1337b0rd/internal/governor/posts"

type Governor struct {
	*posts_governor.PostsGovernor
}

func New() *Governor {
	return &Governor{
		PostsGovernor: new(posts_governor.PostsGovernor),
	}
}
