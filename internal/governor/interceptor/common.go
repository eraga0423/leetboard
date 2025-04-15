package interceptor

import (
	"1337b0rd/internal/config"
	my_redis "1337b0rd/internal/redis"
	"1337b0rd/internal/types/database"
)

type Interceptor struct {
	conf  *config.Config
	db    database.Database
	redis *my_redis.MyRedis
}

func New(conf *config.Config, db database.Database) *Interceptor {
	return &Interceptor{
		db:   db,
		conf: conf,
	}
}
