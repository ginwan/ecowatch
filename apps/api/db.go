package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func connectDB() (*pgxpool.Pool, error) {
	connStr := "postgres://postgres:postgres@localhost:5432/ecowatch"

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
