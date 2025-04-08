package leetboard

import (
	"database/sql"
	"log/slog"
)

type Leetboard struct {
	logger *slog.Logger
	db     *sql.DB
}

func New(
	logger *slog.Logger,
	db *sql.DB,
) *Leetboard {
	return &Leetboard{
		logger: logger,
		db:     db,
	}
}
