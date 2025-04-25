package interceptor

import (
	"1337b0rd/internal/config"
	rickmortyrest "1337b0rd/internal/rick_morty_rest"
	"1337b0rd/internal/types/database"
	redis_types "1337b0rd/internal/types/redis"
)

type Interceptor struct {
	conf        *config.Config
	db          database.Database
	redis       redis_types.TypesRedis
	parseAvatar *rickmortyrest.RickAndMorty
}

func New(conf *config.Config, db database.Database, r redis_types.TypesRedis) *Interceptor {
	return &Interceptor{
		db:    db,
		conf:  conf,
		redis: r,
	}
}
