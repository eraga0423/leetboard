package posts_governor

import (
	"1337b0rd/internal/config"
	"1337b0rd/internal/governor/interceptor"
	miniostorage "1337b0rd/internal/minio_storage"
	"1337b0rd/internal/types/database"
	"1337b0rd/internal/types/rick_morty"
	"bytes"
	"errors"
)

type PostsGovernor struct {
	// logger *log.Logger
	conf        *config.Config
	db          database.Database
	miniostor   miniostorage.MinioStorage /////////////////
	all         allPost
	interceptor interceptor.Interceptor
	ricky       rick_morty.RestRickAndMorty
}

func New(conf *config.Config, db database.Database) *PostsGovernor { //, minio miniostorage.MinioStorage
	return &PostsGovernor{
		// logger: logger,
		conf: conf,
		db:   db,
		// miniostor: minio,
	}
}

func (p *PostsGovernor) checkImageType(data []byte) (string, error) {
	switch {
	case bytes.HasPrefix(data, []byte{0xFF, 0xD8, 0xFF}):
		return "image/jpeg", nil
	case bytes.HasPrefix(data, []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}):
		return "image/png", nil
	default:
		return "", errors.New("invalid image type")
	}
}
