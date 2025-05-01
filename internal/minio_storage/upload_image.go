package miniostorage

import (
	"context"
	"crypto/rand"
	"fmt"
	"log/slog"
	"math/big"
	"mime/multipart"
	"time"

	"1337b0rd/internal/types/storage"

	"github.com/minio/minio-go/v7"
)

type dataImageReq struct {
	bucketName  string
	objectSize  int64
	contentType string
	metadata    multipart.File
}
type dataImageRes struct {
	imageURL string
}

func (m MinioStorage) UploadImage(ctx context.Context, req storage.DataImageReq) (storage.DataImageRes, error) {
	newReq := dataImageReq{
		bucketName:  req.GetBucketName(),
		objectSize:  req.GetObjectSize(),
		contentType: req.GetContentType(),
		metadata:    req.GetMetaData(),
	}

	err := m.client.MakeBucket(ctx, newReq.bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucket := m.client.BucketExists(ctx, newReq.bucketName)
		if errBucket == nil && exists {
			slog.Info("bucket already exists", "bucket", newReq.bucketName)
		} else {
			return nil, fmt.Errorf("failed to create bucket: %v", err)
		}
	}
	newObjectname, err := generateRandomObjectName()
	if err != nil {
		return nil, err
	}
	_, err = m.client.PutObject(ctx, newReq.bucketName, newObjectname, newReq.metadata, newReq.objectSize, minio.PutObjectOptions{
		ContentType: newReq.contentType,
	})
	if err != nil {
		return nil, err
	}
	newURL, err := m.client.PresignedGetObject(ctx, newReq.bucketName, newObjectname, time.Hour*24, nil)
	if err != nil {
		return nil, err
	}

	return &dataImageRes{imageURL: newURL.String()}, nil
}

func generateRandomObjectName() (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	objectNameLength := 16 // Длина имени объекта

	var objectName string
	for i := 0; i < objectNameLength; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random index: %v", err)
		}
		objectName += string(chars[index.Int64()])
	}

	return objectName, nil
}

func (d *dataImageRes) GetImageURL() string {
	return d.imageURL
}
