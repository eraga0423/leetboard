package leetboard

import (
	"1337b0rd/internal/config"
	"database/sql"
	"log/slog"
)

type Leetboard struct {
	conf   *config.PostgresConfig
	logger *slog.Logger
	db     *sql.DB
}

func New(
	conf *config.PostgresConfig,
	logger *slog.Logger,
	db *sql.DB,
) *Leetboard {
	return &Leetboard{
		conf:   conf,
		logger: logger,
		db:     db,
	}
}
