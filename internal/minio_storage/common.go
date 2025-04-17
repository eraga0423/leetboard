package miniostorage

import (
	"1337b0rd/internal/config"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

type MinioStorage struct {
	//conf   *config.Config
	client *minio.Client
}

func NewMinioStorage(conf *config.Config, ctx context.Context) *MinioStorage {
	client := newMinioClient(ctx, conf)
	return &MinioStorage{
		//conf:   conf,
		client: client,
	}
}

func newMinioClient(_ context.Context, conf *config.Config) *minio.Client {
	endpoint := fmt.Sprintf("localhost:%s", conf.Minio.APIPort)

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.Minio.User, conf.Minio.Password, ""),
		Secure: false,
	})
	if err != nil {
		log.Print(err.Error())
		log.Fatal(err)
	}
	return client
}
