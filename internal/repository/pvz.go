package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type pVZRepository struct {
	pool *pgxpool.Pool
}

func NewPVZRepository(pool *pgxpool.Pool) *pVZRepository {
	return &pVZRepository{pool: pool}
}

func (p *pVZRepository) CreatePVZ(ctx context.Context, id, registrationDate, city string) error {
	return nil
}

func (p *pVZRepository) GetPVZList(ctx context.Context) error {
	return nil
}

func (p *pVZRepository) CloseLastReception(ctx context.Context, pvzID string) error {
	return nil
}

func (p *pVZRepository) DeleteLastProduct(ctx context.Context, pvzID string) error {
	return nil
}
