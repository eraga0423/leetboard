package posts_governor

import (
	"1337b0rd/internal/config"
	"1337b0rd/internal/governor/interceptor"
	"1337b0rd/internal/types/database"
	"1337b0rd/internal/types/storage"
)

type PostsGovernor struct {
	// logger *log.Logger
	conf        *config.Config
	db          database.Database
	miniostor   storage.Storage /////////////////
	all         allPost
	interceptor interceptor.Interceptor
}

func New(conf *config.Config, db database.Database, minio storage.Storage) *PostsGovernor { //, minio miniostorage.MinioStorage
	return &PostsGovernor{
		// logger: logger,
		conf:      conf,
		db:        db,
		miniostor: minio,
	}
}
