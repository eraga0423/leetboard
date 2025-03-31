package config

import "os"

type PostgresConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

func NewConfigPostrgres() *PostgresConfig {
	c := &PostgresConfig{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		Name:     os.Getenv("PG_NAME"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
	}
	return c
}
