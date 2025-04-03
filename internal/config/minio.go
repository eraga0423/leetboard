package config

import "os"

type MinioStorage struct {
	User        string
	Password    string
	APIPort     string
	ConsolePort string
}

func NewConfigMinio() *MinioStorage {
	return &MinioStorage{
		User:        os.Getenv("MINIO_ROOT_USER"),
		Password:    os.Getenv("MINIO_ROOT_PASSWORD"),
		APIPort:     os.Getenv("MINIO_API_PORT"),
		ConsolePort: os.Getenv("MINIO_CONSOLE_PORT"),
	}
}
