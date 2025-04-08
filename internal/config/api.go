package config

import "os"

type APIConfig struct {
	Rest     APIRestConfig
	RestRick APIRestRickConfig
}
type APIRestConfig struct {
	Host string
	Port string
}

type APIRestRickConfig struct {
	Host string
	Port string
}

func NewConfigAPI() *APIConfig {
	c := &APIRestConfig{
		Host: os.Getenv("REST_API_HOST"),
		Port: os.Getenv("REST_API_PORT"),
	}
	cr := &APIRestRickConfig{
		Host: os.Getenv("REST_RICK_HOST"),
		Port: os.Getenv("REST_RICK_PORT"),
	}
	return &APIConfig{
		Rest:     *c,
		RestRick: *cr,
	}
}
