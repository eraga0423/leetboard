package governor

import (
	"context"

	redis_types "1337b0rd/internal/types/redis"
	"1337b0rd/internal/types/storage"

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

func (g *Governor) ConfigGov(
	_ context.Context,
	conf *config.Config,
	db database.Database,
	r redis_types.TypesRedis,
	minio storage.Storage,
) {
	*g.Interceptor = *interceptor.New(conf, db, r)
	*g.PostsGovernor = *posts_governor.New(conf, db, minio)
}
