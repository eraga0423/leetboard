package ports

import "1337b0rd/internal/core"

type PostRepository interface {
	Save(post *core.Post) (*core.Post, error)
}
