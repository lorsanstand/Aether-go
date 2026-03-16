package database

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type config interface {
	GetUrlPostgres() string
}

func NewPostgresDB(cfg config) (*sql.DB, error) {
	conn, err := sql.Open("pgx", cfg.GetUrlPostgres())
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

func NewPostgresPGX(ctx context.Context, cfg config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, cfg.GetUrlPostgres())
	if err != nil {
		return nil, err
	}
	return pool, nil
}
