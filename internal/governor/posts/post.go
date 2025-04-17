package posts_governor

import (
	"1337b0rd/internal/config"
	"1337b0rd/internal/governor/interceptor"
	miniostorage "1337b0rd/internal/minio_storage"
	"1337b0rd/internal/types/database"
	"1337b0rd/internal/types/rick_morty"
)

type PostsGovernor struct {
	// logger *log.Logger
	conf        *config.Config
	db          database.Database
	miniostor   miniostorage.MinioStorage /////////////////
	all         allPost
	interceptor interceptor.Interceptor
	ricky       rick_morty.RestRickAndMorty
}

func New(conf *config.Config, db database.Database) *PostsGovernor { //, minio miniostorage.MinioStorage
	return &PostsGovernor{
		// logger: logger,
		conf: conf,
		db:   db,
		// miniostor: minio,
	}
}
