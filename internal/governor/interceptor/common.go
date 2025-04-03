package interceptor

import "1337b0rd/internal/config"

type Interceptor struct {
	conf *config.Config
}

func New(conf *config.Config) *Interceptor {
	return &Interceptor{
		conf: conf,
	}
}
