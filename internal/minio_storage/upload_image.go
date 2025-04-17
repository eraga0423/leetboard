package miniostorage

import (
	"1337b0rd/internal/types/storage"
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"time"
)

type dataImageReq struct {
	bucketName  string
	objectName  string
	objectSize  int64
	contentType string
	metadata    io.Reader
}
type dataImageRes struct {
	imageURL string
}

func (m MinioStorage) UploadImage(ctx context.Context, req storage.DataImageReq) (storage.DataImageRes, error) {
	newReq := dataImageReq{
		bucketName:  req.GetBucketName(),
		objectName:  req.GetObjectName(),
		objectSize:  req.GetObjectSize(),
		contentType: req.GetContentType(),
		metadata:    req.GetMetadata(),
	}

	_, err := m.client.PutObject(ctx, newReq.bucketName, newReq.objectName, newReq.metadata, newReq.objectSize, minio.PutObjectOptions{
		ContentType: newReq.contentType,
	})
	if err != nil {
		return nil, err
	}
	newURL, err := m.client.PresignedGetObject(ctx, newReq.bucketName, newReq.objectName, time.Hour*24, nil)
	if err != nil {
		return nil, err
	}

	return &dataImageRes{imageURL: newURL.String()}, nil
}

func (d *dataImageRes) GetImageURL() string {
	return d.imageURL
}
