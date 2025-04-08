package governor

import (
	"context"

	"1337b0rd/internal/config"
	"1337b0rd/internal/governor/interceptor"
	posts_governor "1337b0rd/internal/governor/posts"
	"1337b0rd/internal/types/database"
)

type Governor struct {
	*posts_governor.PostsGovernor
	*interceptor.Interceptor
}

func New() *Governor {
	return &Governor{
		PostsGovernor: new(posts_governor.PostsGovernor),
		Interceptor:   new(interceptor.Interceptor),
	}
}

func (g *Governor) ConfigGov(ctx context.Context, conf *config.Config, db database.Database) {
	*g.Interceptor = *interceptor.New(conf)
	*g.PostsGovernor = *posts_governor.New(conf, db)
	
}
