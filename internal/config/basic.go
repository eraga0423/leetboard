package config

type Config struct {
	API      APIConfig
	Postgres PostgresConfig
	Minio    MinioStorage
}

func NewConfig() *Config {
	return &Config{
		API:      *NewConfigAPI(),
		Postgres: *NewConfigPostrgres(),
		Minio:    *NewConfigMinio(),
	}
}
