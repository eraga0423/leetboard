package miniostorage

import (
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func (m MinioStorage) newMinioClient() *minio.Client {
	endpoint := fmt.Sprintf("localhost:%s", m.conf.APIPort)
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.conf.User, m.conf.Password, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatal(err)
	}
	return client
}

