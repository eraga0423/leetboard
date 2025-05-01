package config

import "os"

type MinioStorage struct {
	User        string
	Password    string
	APIPort     string
	ConsolePort string
	Host        string
	PublicHost  string
}

func NewConfigMinio() *MinioStorage {
	return &MinioStorage{
		PublicHost:  os.Getenv("MINIO_PUBLIC_HOST"),
		User:        os.Getenv("MINIO_ROOT_USER"),
		Password:    os.Getenv("MINIO_ROOT_PASSWORD"),
		APIPort:     os.Getenv("MINIO_API_PORT"),
		ConsolePort: os.Getenv("MINIO_CONSOLE_PORT"),
		Host:        os.Getenv("MINIO_HOST"),
	}
}
