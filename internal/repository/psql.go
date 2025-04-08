package repository

import (
	"context"

	"github.com/Wammero/PVZ-service/internal/config.go"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func New(cfg *config.DatabaseConfig) (*Repository, error) {
	pool, err := pgxpool.Connect(context.Background(), cfg.GetConnStr())
	if err != nil {
		return nil, err
	}

	return &Repository{pool: pool}, nil
}

func (repo *Repository) Close() {
	if repo.pool != nil {
		repo.pool.Close()
	}
}
