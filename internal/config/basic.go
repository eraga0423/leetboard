package config

type Config struct {
	API      APIConfig
	Postgres PostgresConfig
}

func NewConfig() *Config {
	return &Config{
		API:      *NewConfigAPI(),
		Postgres: *NewConfigPostrgres(),
	}
}
