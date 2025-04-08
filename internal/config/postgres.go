package config

import (
	"os"
	"strconv"
)

type PostgresConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

func NewConfigPostrgres() *PostgresConfig {
	port, _ := strconv.Atoi("PG_PORT")
	c := &PostgresConfig{
		Host:     os.Getenv("PG_HOST"),
		Port:     port,
		Name:     os.Getenv("PG_NAME"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
	}
	return c
}
