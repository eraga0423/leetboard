package my_redis

import (
	"fmt"
	"log/slog"

	"1337b0rd/internal/config"
	"1337b0rd/internal/types/controller"

	"github.com/redis/go-redis/v9"
)

type MyRedis struct {
	ctrl      controller.Controller
	config    *config.Config
	newClient *redis.Client
}

func NewMyRedis(ctrl controller.Controller, config *config.Config, log *slog.Logger) *MyRedis {
	return &MyRedis{
		ctrl:   ctrl,
		config: config,
		newClient: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		}),
	}
}

//func (m MyRedis) newClient() *redis.Client {
//	return redis.NewClient(&redis.Options{
//		Addr: fmt.Sprintf("%s:%s", m.config.Redis.Host, m.config.Redis.Port),
//	})
//}
