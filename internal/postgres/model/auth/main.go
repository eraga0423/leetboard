package auth

import (
	"database/sql"
	"log/slog"
)

type Auth struct {
	logger *slog.Logger
	db     *sql.DB
}

func New(
	logger *slog.Logger,
	db *sql.DB,
) *Auth {
	return &Auth{
		logger: logger,
		db:     db,
	}
}
