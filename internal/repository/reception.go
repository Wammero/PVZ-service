package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type receptionRepositor struct {
	pool *pgxpool.Pool
}

func NewReceptionRepository(pool *pgxpool.Pool) *receptionRepositor {
	return &receptionRepositor{pool: pool}
}

func (r *receptionRepositor) CreateReception(ctx context.Context, pvzId string) error {
	return nil
}
