package posts_governor

import (
	"1337b0rd/internal/config"
	"1337b0rd/internal/types/database"
)

type PostsGovernor struct {
	// logger *log.Logger
	conf *config.Config
	db   database.Database
}

func New(conf *config.Config, db database.Database) *PostsGovernor {
	return &PostsGovernor{
		// logger: logger,
		conf: conf,
		db:   db,
	}
}
