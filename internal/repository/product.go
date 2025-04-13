package repository

import (
	"context"
	"fmt"

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

func (p *productRepository) AddProduct(ctx context.Context, tx pgx.Tx, productType string, pvzId string) (string, string, string, error) {
	var receptionID string
	checkReceptionQuery := `
		SELECT reception_id FROM receptions
		WHERE pvz_id = $1 AND status = 'in_progress'
		ORDER BY reception_time DESC
		LIMIT 1;
	`
	err := tx.QueryRow(ctx, checkReceptionQuery, pvzId).Scan(&receptionID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", "", "", fmt.Errorf("нет открытой приёмки для ПВЗ %s", pvzId)
		}
		return "", "", "", fmt.Errorf("ошибка при поиске активной приёмки: %w", err)
	}

	var productID, receptionTime string
	insertProductQuery := `
		INSERT INTO products (type)
		VALUES ($1)
		RETURNING product_id::text, reception_time::text;
	`
	err = tx.QueryRow(ctx, insertProductQuery, productType).Scan(&productID, &receptionTime)
	if err != nil {
		return "", "", "", fmt.Errorf("не удалось добавить товар: %w", err)
	}

	linkQuery := `
		INSERT INTO reception_products (reception_id, product_id)
		VALUES ($1, $2);
	`
	_, err = tx.Exec(ctx, linkQuery, receptionID, productID)
	if err != nil {
		return "", "", "", fmt.Errorf("не удалось привязать товар к приёмке: %w", err)
	}

	return productID, receptionTime, receptionID, nil
}
