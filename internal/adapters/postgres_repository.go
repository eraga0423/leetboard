package adapters

import (
	"database/sql"

	"1337b0rd/internal/core"
	"1337b0rd/internal/ports"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.PostRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Save(post *core.Post) (*core.Post, error) {
	query := "INSERT INTO posts (content, image) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, post.Content, post.Image).Scan(&post.ID)
	if err != nil {
		return nil, err
	}
	return post, nil
}
