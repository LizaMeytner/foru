package postgres

import (
	"database/sql"

	"github.com/LizaMeytner/foru/forum-service/config"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg config.PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
