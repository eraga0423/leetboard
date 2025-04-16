package main

import (
	"context"
	"log"
	"log/slog"

	"1337b0rd/internal/config"
	"1337b0rd/internal/governor"
	postgres "1337b0rd/internal/postgres"
	"1337b0rd/internal/rest"

	_ "github.com/lib/pq"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	gov := governor.New()
	err := gov.Interceptor.FetchAndCacheAvatar(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	r := rest.New(gov)
	conf := config.NewConfig()
	p, err := postgres.New(&conf.Postgres, nil)
	if err != nil {
		slog.Any("failed start database", "postgres")
		panic(err)
	}
	go func(ctx context.Context, cancelFunc context.CancelFunc, apiConfig config.APIConfig) {
		err := r.Start(ctx, &apiConfig)
		if err != nil {
			log.Fatal(err)
		}

		cancelFunc()
	}(ctx, cancel, conf.API)
	gov.ConfigGov(ctx, conf, p)
	<-ctx.Done()
}
