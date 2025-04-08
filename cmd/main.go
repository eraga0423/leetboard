package main

import (
	"context"
	"log"
	"log/slog"

	"1337b0rd/internal/config"
	"1337b0rd/internal/governor"
	posgres "1337b0rd/internal/postgres"
	"1337b0rd/internal/rest"

	_ "github.com/lib/pq"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	gov := governor.New()

	r := rest.New(gov)
	conf := config.NewConfig()
	p, err := posgres.NewPosgres(&conf.Postgres)
	if err != nil {
		slog.Any("failed start database", "postgres")
		panic(err)
	}

	go func(ctx context.Context, cancelFunc context.CancelFunc) {
		err := r.Start(ctx)
		if err != nil {
			log.Fatal(err)
		}
		cancelFunc()
	}(ctx, cancel)
	gov.ConfigGov(ctx, conf, p)

}
