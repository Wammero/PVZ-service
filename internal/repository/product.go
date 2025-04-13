package repository

import (
	"context"
	"fmt"

	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepository struct {
	pool *pgxpool.Pool
}

func NewProductRepository(pool *pgxpool.Pool) *productRepository {
	return &productRepository{pool: pool}
}

func (r *productRepository) Pool() *pgxpool.Pool {
	return r.pool
}

func (p *productRepository) AddProduct(ctx context.Context, tx pgx.Tx, productType string, pvzId string) (*model.Product, error) {
	query := `
		WITH active_reception AS (
			SELECT reception_id
			FROM receptions
			WHERE pvz_id = $1 AND status = 'in_progress'
			ORDER BY reception_time DESC
			LIMIT 1
		),
		inserted_product AS (
			INSERT INTO products (type)
			VALUES ($2)
			RETURNING product_id, reception_time
		),
		link_product AS (
			INSERT INTO reception_products (reception_id, product_id)
			SELECT ar.reception_id, ip.product_id
			FROM active_reception ar, inserted_product ip
		)
		SELECT ip.product_id, ip.reception_time, ar.reception_id
		FROM inserted_product ip, active_reception ar;
	`

	var product model.Product
	err := tx.QueryRow(ctx, query, pvzId, productType).Scan(
		&product.ID,
		&product.Date,
		&product.ReceptionID,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("нет активной приёмки для ПВЗ %s", pvzId)
		}
		return nil, fmt.Errorf("ошибка при добавлении товара: %w", err)
	}

	product.Type = productType

	return &product, nil
}
