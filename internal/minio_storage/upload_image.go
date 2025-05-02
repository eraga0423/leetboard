package miniostorage

import (
	"context"
	"crypto/rand"
	"fmt"
	"log/slog"
	"math/big"
	"mime/multipart"

	"1337b0rd/internal/types/storage"

	"github.com/minio/minio-go/v7"
)

type dataImageReq struct {
	bucketName  string
	objectSize  int64
	contentType string
	metadata    multipart.File
	objectName  string
}
type dataImageRes struct {
	imageURL string
}
type newParseResp struct {
	objectName string
	newURL     string
}

func (m *MinioStorage) UploadImage(ctx context.Context, req storage.DataImageReq) error {
	newReq := dataImageReq{
		bucketName:  req.GetBucketName(),
		objectSize:  req.GetObjectSize(),
		contentType: req.GetContentType(),
		metadata:    req.GetMetaData(),
		objectName:  req.GetObjectName(),
	}
	err := m.client.MakeBucket(ctx, newReq.bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucket := m.client.BucketExists(ctx, newReq.bucketName)
		if errBucket == nil && exists {
			slog.Info("bucket already exists", "bucket", newReq.bucketName)
		} else {
			return fmt.Errorf("failed to create bucket: %v", err)
		}
	}
	_, err = m.client.PutObject(ctx, newReq.bucketName, newReq.objectName, newReq.metadata, newReq.objectSize, minio.PutObjectOptions{
		ContentType: newReq.contentType,
	})
	if err != nil {
		return fmt.Errorf("failed to create bucket: %v", err)
	}
	policy := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [
		  {
			"Effect": "Allow",
			"Principal": {"AWS": ["*"]},
			"Action": ["s3:GetObject"],
			"Resource": ["arn:aws:s3:::%s/*"]
		  }
		]
	  }`, newReq.bucketName)

	err = m.client.SetBucketPolicy(ctx, newReq.bucketName, policy)
	if err != nil {
		return fmt.Errorf("bucket policy %v", err)
	}
	return nil
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

func (m *MinioStorage) ParseURL(ctx context.Context, req storage.DataImageReq) (storage.DataImageRes, error) {
	newReq := dataImageReq{bucketName: req.GetBucketName()}
	newObjectname, err := generateRandomObjectName()
	if err != nil {
		return nil, err
	}
	newURL := fmt.Sprintf("http://%s/%s/%s", m.conf.Minio.PublicHost, newReq.bucketName, newObjectname)

	return &newParseResp{
		objectName: newObjectname,
		newURL:     newURL,
	}, nil
}

func (u *newParseResp) GetImageURL() string {
	return u.newURL
}

func (u *newParseResp) GetNewObjectName() string {
	return u.objectName
}
