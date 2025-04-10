package repository

import (
	"context"

	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepository struct {
	pool *pgxpool.Pool
}

func NewProductRepository(pool *pgxpool.Pool) *productRepository {
	return &productRepository{pool: pool}
}

func (p *productRepository) AddProduct(ctx context.Context, productType model.ProductType, pvzId string) error {
	return nil
}
