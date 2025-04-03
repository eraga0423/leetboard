package miniostorage

import (
	"bytes"
	"context"
	"time"

	"github.com/minio/minio-go/v7"
)

func (m MinioStorage) UploadImage(bucketName, objectName, contentType string, data []byte) (string, error) {
	client := m.newMinioClient()
	ctx := context.Background()
	reader := bytes.NewReader(data)

	_, err := client.PutObject(ctx, bucketName, objectName, reader, int64(len(data)), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", err
	}
	newURL, err := client.PresignedGetObject(ctx, bucketName, objectName, time.Hour*24, nil)
	if err != nil {
		return "", err
	}
	return newURL.String(), nil
}
