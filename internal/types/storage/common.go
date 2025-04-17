package storage

import "context"

type Storage interface {
	UploadImage(context.Context, DataImageReq) (DataImageRes, error)
}
