package leetboard

import (
	"1337b0rd/internal/config"
	"1337b0rd/internal/models"
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

func (m models.Post) CreatePost() models.Post {
	m.PostID = 5
	return m
}
