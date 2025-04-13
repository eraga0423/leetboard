package config

import "os"

type RedisConfig struct {
	Host string
	Port string
}

func NewRedisConfig() *RedisConfig {
	r := &RedisConfig{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
	}
	return r

}
