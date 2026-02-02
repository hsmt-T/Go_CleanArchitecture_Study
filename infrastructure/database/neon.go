package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewNeon(ctx context.Context) (*pgxpool.Pool, error) {
	dsn := os.Getenv("NEON_DATABASE_URL")

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	return pgxpool.NewWithConfig(ctx, config)
}
