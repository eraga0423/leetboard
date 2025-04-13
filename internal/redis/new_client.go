package my_redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

func (m MyRedis) NewClient() {
	redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", m.config.Redis.Host, m.config.Redis.Port),
	})

}
