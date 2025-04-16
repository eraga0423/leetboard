package interceptor

import (
	"1337b0rd/internal/config"
	my_redis "1337b0rd/internal/redis"
	rickmortyrest "1337b0rd/internal/rick_morty_rest"
	"1337b0rd/internal/types/database"
)

type Interceptor struct {
	conf        *config.Config
	db          database.Database
	redis       *my_redis.MyRedis
	parseAvatar *rickmortyrest.RickAndMorty
}

func New(conf *config.Config, db database.Database) *Interceptor {
	return &Interceptor{
		db:   db,
		conf: conf,
	}
}
