package storage

import "io"

type DataImageReq interface {
	GetBucketName() string
	GetObjectName() string
	GetObjectSize() int64
	GetContentType() string
	GetMetaData() io.Reader
}
type DataImageRes interface {
	GetImageURL() string
}
