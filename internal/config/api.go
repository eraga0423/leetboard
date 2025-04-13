package config

import "os"

type APIConfig struct {
	Rest APIRestConfig
}
type APIRestConfig struct {
	Host string
	Port string
}

func NewConfigAPI() *APIConfig {
	c := &APIRestConfig{
		Host: os.Getenv("REST_API_HOST"),
		Port: os.Getenv("REST_API_PORT"),
	}

	return &APIConfig{
		Rest: *c,
	}
}
