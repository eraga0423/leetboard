package auth

import (
	"1337b0rd/internal/config"
	"database/sql"
	"log/slog"
)

type Auth struct {
	conf   *config.PostgresConfig
	logger *slog.Logger
	db     *sql.DB
}

func New(
	conf *config.PostgresConfig,
	logger *slog.Logger,
	db *sql.DB,
) *Auth {
	return &Auth{
		conf:   conf,
		logger: logger,
		db:     db,
	}
}
