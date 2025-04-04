package posts_governor

import (
	"1337b0rd/internal/config"
	miniostorage "1337b0rd/internal/minio_storage"
	"1337b0rd/internal/types/database"
)

type PostsGovernor struct {
	// logger *log.Logger
	conf      *config.Config
	db        database.Database
	miniostor miniostorage.MinioStorage /////////////////
}

func New(conf *config.Config, db database.Database) *PostsGovernor { //, minio miniostorage.MinioStorage
	return &PostsGovernor{
		// logger: logger,
		conf:      conf,
		db:        db,
		// miniostor: minio,
	}
}
