package posgres

import (
	"1337b0rd/internal/config"
	"1337b0rd/internal/posgres/auth"
	"database/sql"
	"fmt"
)

type Postgres struct {
	*auth.Auth
}

func NewDB(conf *config.PostgresConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Name,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
