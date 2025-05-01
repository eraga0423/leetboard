package main

import (
	"context"
	"io"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"sync"
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
	var wg sync.WaitGroup
	// логирование начало
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatalf("Не удалось открыть лог-файл: %v", err)
	}
	defer file.Close()

	// Создаём мульти-вывод: в stdout и в файл одновременно
	multiWriter := io.MultiWriter(os.Stdout, file)
	handler := slog.NewTextHandler(multiWriter, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	logger := slog.New(handler)
	// логирование конец

	// реализуется тайм тиккер - начало
	wg.Add(1)
	go func(ctx context.Context, cancelFunc context.CancelFunc, logger *slog.Logger, wg *sync.WaitGroup) {
		ticker := time.NewTicker(time.Minute * 5)
		defer ticker.Stop()
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				err := gov.Interceptor.BackupAvatars(ctx)
				if err != nil {
					logger.ErrorContext(ctx, "backup avatars could not update", "error", err)
					cancelFunc()
					return
				}
				logger.Info("Auto save backup work succesfully")
			}
		}
	}(ctx, cancel, logger.With(slog.String("service", "backup avatars")), &wg)
	// реализуется тайм тиккер - конец

	// rickmortyrest.NewRickAndMorty()
	r := rest.New(gov)
	conf := config.NewConfig()
	ms := miniostorage.NewMinioStorage(conf, ctx)

	// postgres начало
	p, err := postgres.New(&conf.Postgres, logger.With(slog.String("service", "postgre")))
	if err != nil {
		logger.ErrorContext(ctx, "failed to start postgre", slog.Any("error", err))
	}
	// postgres конец

	// rest подключение
	wg.Add(1)
	go func(ctx context.Context, cancelFunc context.CancelFunc, apiConfig config.APIConfig, logger *slog.Logger, wg *sync.WaitGroup) {
		err := r.Start(ctx, cancelFunc, &apiConfig, logger.With(slog.String("service", "rest")))
		if err != nil {
			logger.ErrorContext(ctx, "failed to start rest", slog.Any("error", err))
		}

		cancelFunc()
		wg.Done()
	}(ctx, cancel, conf.API, logger, &wg)
	// rest конец

	// redis подключение
	myRedis := my_redis.NewMyRedis(gov, conf, logger.With(slog.String("service", "redis")))
	// redis конец

	// формирование конструктура начало
	gov.ConfigGov(ctx, conf, p, myRedis, ms)
	// формирование конструктура конец

	// скачивание аватаров и присваивание в базу данных начало
	err = gov.Interceptor.FetchAndCacheAvatar(ctx, logger.With(slog.String("service", "fetch and cache avatar")))
	if err != nil {
		logger.ErrorContext(ctx, "failed to start configGov", slog.Any("error", err))
		return
	}
	// скачивание аватаров и присваивание в базу данных конец

	// gracefullshutdown начало
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
	// gracefullshutdown конец

	// ждет функцию cancel
	<-ctx.Done()
	// ждет чтобы все go rutine завершились
	wg.Wait()
}
