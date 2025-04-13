package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type receptionRepositor struct {
	pool *pgxpool.Pool
}

func NewReceptionRepository(pool *pgxpool.Pool) *receptionRepositor {
	return &receptionRepositor{pool: pool}
}

func (r *receptionRepositor) Pool() *pgxpool.Pool {
	return r.pool
}

func (r *receptionRepositor) CreateReception(ctx context.Context, tx pgx.Tx, pvzId string) (string, string, error) {
	checkQuery := `
		SELECT 1 FROM receptions
		WHERE pvz_id = $1 AND status = 'in_progress'
		LIMIT 1;
	`

	var exists int
	err := tx.QueryRow(ctx, checkQuery, pvzId).Scan(&exists)
	if err != nil && err.Error() != "no rows in result set" {
		return "", "", fmt.Errorf("ошибка при проверке существующей приёмки: %v", err)
	}
	if exists == 1 {
		return "", "", fmt.Errorf("уже есть приёмка со статусом 'in_progress' для ПВЗ %s", pvzId)
	}

	insertQuery := `
		INSERT INTO receptions (pvz_id, status)
		VALUES ($1, 'in_progress')
		RETURNING reception_id::text, reception_time::text;
	`

	var receptionId, dateTime string
	err = tx.QueryRow(ctx, insertQuery, pvzId).Scan(&receptionId, &dateTime)
	if err != nil {
		return "", "", fmt.Errorf("не удалось создать приёмку: %v", err)
	}

	return receptionId, dateTime, nil
}
