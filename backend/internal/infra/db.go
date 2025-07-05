package infra

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"restservice/internal/exceptions"
)

type PostgresConnect struct {
	Pool *pgxpool.Pool
}

func NewPostgresConnect(url string) *PostgresConnect {
	conn, err := pgxpool.New(context.Background(), url)
	if err != nil {
		exceptions.HandleError(&exceptions.CustomException{Field: "DB", Message: "failed to connect"})
	}
	return &PostgresConnect{Pool: conn}
}

func (pg *PostgresConnect) Close() {
	if pg.Pool != nil {
		pg.Pool.Close()
	}
}
