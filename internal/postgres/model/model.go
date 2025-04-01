package model

import (
	"1337b0rd/internal/config"
	"1337b0rd/internal/postgres/auth"
	"1337b0rd/internal/postgres/leetboard"
	"database/sql"
	"log/slog"
)

type Model struct {
	*auth.Auth
	*leetboard.Leetboard
}

func New(
	conf *config.PostgresConfig,
	logger *slog.Logger,
	db *sql.DB,
) *Model {
	return &Model{
		Auth:      auth.New(conf, logger.With(slog.String("component", "auth")), db),
		Leetboard: leetboard.New(conf, logger.With(slog.String("component", "leetboard")), db),
	}
}
