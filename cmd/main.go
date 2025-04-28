package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"time"

	miniostorage "1337b0rd/internal/minio_storage"
	my_redis "1337b0rd/internal/redis"

	"1337b0rd/internal/config"
	"1337b0rd/internal/governor"
	postgres "1337b0rd/internal/postgres"
	"1337b0rd/internal/rest"

	_ "github.com/lib/pq"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	gov := governor.New()

	go func(ctx context.Context) {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				err := gov.Interceptor.BackupAvatars(ctx)
				if err != nil {
					log.Print(err)
					return
				}
				log.Print("auto save backup succesfully")
			}
		}
	}(ctx)
	////////////////////rickmortyrest.NewRickAndMorty()
	r := rest.New(gov)
	conf := config.NewConfig()
	ms := miniostorage.NewMinioStorage(conf, ctx)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)) ///////////
	p, err := postgres.New(&conf.Postgres, logger)
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
	myRedis := my_redis.NewMyRedis(gov, conf)
	gov.ConfigGov(ctx, conf, p, myRedis, ms)
	err = gov.Interceptor.FetchAndCacheAvatar(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	go func(cancelFunc context.CancelFunc) {
		shutdown := make(chan os.Signal, 1)   // Create channel to signify s signal being sent
		signal.Notify(shutdown, os.Interrupt) // When an interrupt is sent, notify the channel

		sig := <-shutdown
		slog.Any("signal", sig)
		err := gov.Interceptor.BackupAvatars(ctx)
		if err != nil {
			log.Print(err)
			return
		}
		log.Print("auto save backup succesfully")
		cancelFunc()
	}(cancel)

	<-ctx.Done()
}
