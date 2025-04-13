package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type pVZRepository struct {
	pool *pgxpool.Pool
}

func NewPVZRepository(pool *pgxpool.Pool) *pVZRepository {
	return &pVZRepository{pool: pool}
}

func (p *pVZRepository) CreatePVZ(ctx context.Context, id, city string, regDate time.Time, creator sql.NullInt64) error {
	query := `
		INSERT INTO pvz (id, creator_id, registration_date, city)
		VALUES ($1, $2, $3, $4);`

	_, err := p.pool.Exec(ctx, query, id, creator, regDate, city)
	if err != nil {
		if pgx.ErrNoRows == err {
			return fmt.Errorf("пункт приёма заказов с таким id уже существует")
		}
		return fmt.Errorf("не удалось создать ПВЗ: %v", err)
	}

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
