package my_redis

import (
	"1337b0rd/internal/config"
	"1337b0rd/internal/types/controller"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type MyRedis struct {
	ctrl   controller.Controller
	config *config.Config
}

func NewMyRedis(ctrl controller.Controller, config *config.Config) *MyRedis {
	return &MyRedis{ctrl: ctrl, config: config}
}

func (m MyRedis) newClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", m.config.Redis.Host, m.config.Redis.Port),
	})
}
