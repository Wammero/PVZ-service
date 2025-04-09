package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	Authorization
	PVZ
	Reception
	Product
	pool *pgxpool.Pool
}

func New(connstr string) (*Repository, error) {
	pool, err := pgxpool.Connect(context.Background(), connstr)
	if err != nil {
		return nil, err
	}

	return &Repository{
		Authorization: NewAuthRepository(pool),
		PVZ:           NewPVZRepository(pool),
		Reception:     NewReceptionRepository(pool),
		Product:       NewProductRepository(pool),
		pool:          pool,
	}, nil
}

func (r *Repository) Close() {
	if r.pool != nil {
		r.pool.Close()
	}
}
