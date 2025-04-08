package model

import (
	"database/sql"
	"log/slog"

	"1337b0rd/internal/postgres/model/auth"
	"1337b0rd/internal/postgres/model/leetboard"
)

type Model struct {
	*auth.Auth
	*leetboard.Leetboard
}

func New(
	logger *slog.Logger,
	db *sql.DB,
) *Model {
	return &Model{
		Auth:      auth.New(logger.With(), db),
		Leetboard: leetboard.New(logger.With(), db),
	}
}
