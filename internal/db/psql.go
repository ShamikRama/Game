package db

import (
	"Game/internal/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func New(cfg config.Config) (*sql.DB, error) {
	const op = "storage.psql.New"

	// Используем формат postgres://
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: failed to ping database: %w", op, err)
	}

	return db, nil
}
