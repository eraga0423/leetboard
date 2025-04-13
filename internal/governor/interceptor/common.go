package interceptor

import (
	"1337b0rd/internal/config"
	"1337b0rd/internal/types/database"
)

type Interceptor struct {
	conf *config.Config
	db database.Database
}

func New(conf *config.Config, db database.Database) *Interceptor {
	return &Interceptor{
		db:db,
		conf: conf,
	}
}
