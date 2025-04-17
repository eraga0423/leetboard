package storage

import (
	"mime/multipart"
)

type DataImageReq interface {
	GetBucketName() string
	GetObjectName() string
	GetObjectSize() int64
	GetContentType() string
	GetMetaData() multipart.File
}
type DataImageRes interface {
	GetImageURL() string
}
