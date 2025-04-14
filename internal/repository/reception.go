package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type receptionRepository struct {
	pool *pgxpool.Pool
}

func NewReceptionRepository(pool *pgxpool.Pool) *receptionRepository {
	return &receptionRepository{pool: pool}
}

func (r *receptionRepository) Pool() *pgxpool.Pool {
	return r.pool
}

func (r *receptionRepository) CreateReception(ctx context.Context, tx pgx.Tx, pvzId string) (*model.Reception, error) {
	query := `
		WITH insert_reception AS (
			INSERT INTO receptions (pvz_id, status)
			SELECT $1, 'in_progress'
			WHERE NOT EXISTS (
				SELECT 1 FROM receptions WHERE pvz_id = $1 AND status = 'in_progress'
			)
			RETURNING reception_id, reception_time
		)
		SELECT reception_id, reception_time FROM insert_reception;
	`

	var receptionID string
	var receptionTime time.Time
	err := tx.QueryRow(ctx, query, pvzId).Scan(&receptionID, &receptionTime)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, fmt.Errorf("уже есть приёмка со статусом 'in_progress' для ПВЗ %s", pvzId)
		}
		return nil, fmt.Errorf("ошибка при создании приёмки: %v", err)
	}

	return &model.Reception{
		ID:     receptionID,
		Date:   receptionTime,
		PVZID:  pvzId,
		Status: "in_progress",
	}, nil
}
