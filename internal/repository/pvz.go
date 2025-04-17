package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type pvzRepository struct {
	pool *pgxpool.Pool
}

func NewPVZRepository(pool *pgxpool.Pool) *pvzRepository {
	return &pvzRepository{pool: pool}
}

func (r *pvzRepository) Pool() *pgxpool.Pool {
	return r.pool
}

func (p *pvzRepository) CreatePVZ(ctx context.Context, id, city string, registrationDate time.Time, creator sql.NullString) error {
	query := `
		INSERT INTO pvz (id, creator_id, registration_date, city)
		VALUES ($1, $2, $3, $4);`

	_, err := p.pool.Exec(ctx, query, id, creator, registrationDate, city)
	if err != nil {
		if pgx.ErrNoRows == err {
			return fmt.Errorf("пункт приёма заказов с таким id уже существует")
		}
		return fmt.Errorf("не удалось создать ПВЗ: %v", err)
	}

	return nil
}

func (p *pvzRepository) GetPVZList(ctx context.Context, tx pgx.Tx, startDate, endDate time.Time, page, limit int) ([]model.PVZWithReceptions, error) {
	offset := (page - 1) * limit
	query := `
		WITH selected_receptions AS (
			SELECT r.reception_id, r.reception_time, r.status, r.pvz_id
			FROM receptions r
			WHERE r.reception_time BETWEEN $1 AND $2
			AND r.status != 'close'
		),
		selected_pvz AS (
			SELECT DISTINCT p.pvz, p.id, p.city, p.registration_date
			FROM pvz p
			JOIN selected_receptions r ON r.pvz_id = p.id
			ORDER BY p.registration_date
			OFFSET $3 LIMIT $4
		)
		SELECT 
			p.pvz, p.id, p.city, p.registration_date,
			r.reception_id, r.reception_time, r.status,
			pr.product_id, pr.reception_time, pr.type
		FROM selected_pvz p
		JOIN selected_receptions r ON r.pvz_id = p.id
		LEFT JOIN reception_products rp ON rp.reception_id = r.reception_id AND rp.is_active = TRUE
		LEFT JOIN products pr ON pr.product_id = rp.product_id AND pr.is_active = TRUE
		ORDER BY p.registration_date, r.reception_time DESC

	`

	rows, err := tx.Query(ctx, query, startDate, endDate, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query pvz list: %w", err)
	}
	defer rows.Close()

	result := make(map[string]*model.PVZWithReceptions)
	for rows.Next() {
		var pvzID int
		var pvzUUID string
		var city string
		var regDate time.Time
		var receptionID, status string
		var receptionTime time.Time
		var productID, productType sql.NullString
		var productTime sql.NullTime

		err := rows.Scan(&pvzID, &pvzUUID, &city, &regDate,
			&receptionID, &receptionTime, &status,
			&productID, &productTime, &productType)
		if err != nil {
			return nil, fmt.Errorf("failed to scan pvz row: %w", err)
		}

		pvz, ok := result[pvzUUID]
		if !ok {
			pvz = &model.PVZWithReceptions{
				PVZ: model.PVZ{
					ID:               pvzUUID,
					City:             city,
					RegistrationDate: regDate,
				},
			}
			result[pvzUUID] = pvz
		}

		var foundReception *model.ReceptionWithProducts
		for i := range pvz.Receptions {
			if pvz.Receptions[i].Reception.ID == receptionID {
				foundReception = &pvz.Receptions[i]
				break
			}
		}
		if foundReception == nil {
			pvz.Receptions = append(pvz.Receptions, model.ReceptionWithProducts{
				Reception: model.Reception{
					ID:     receptionID,
					Status: status,
					Date:   receptionTime,
					PVZID:  pvzUUID,
				},
			})
			foundReception = &pvz.Receptions[len(pvz.Receptions)-1]
		}

		if productID.Valid {
			foundReception.Products = append(foundReception.Products, model.Product{
				ID:          productID.String,
				Date:        productTime.Time,
				Type:        productType.String,
				ReceptionID: receptionID,
			})
		}
	}

	var pvzList []model.PVZWithReceptions
	for _, v := range result {
		pvzList = append(pvzList, *v)
	}
	return pvzList, nil
}

func (p *pvzRepository) CloseLastReception(ctx context.Context, tx pgx.Tx, pvzID string) (*model.Reception, error) {
	var reception model.Reception

	query := `
		WITH updated AS (
			UPDATE receptions
			SET status = 'close'
			WHERE pvz_id = $1 AND status = 'in_progress'
			RETURNING reception_id, reception_time, pvz_id, status
		),
		deactivated_rp AS (
			UPDATE reception_products
			SET is_active = FALSE
			WHERE reception_id IN (SELECT reception_id FROM updated)
			RETURNING product_id
		),
		_ AS (
			UPDATE products
			SET is_active = FALSE
			WHERE product_id IN (SELECT product_id FROM deactivated_rp)
		)
		SELECT reception_id, reception_time, pvz_id, status FROM updated;
	`

	err := tx.QueryRow(ctx, query, pvzID).Scan(&reception.ID, &reception.Date, &reception.PVZID, &reception.Status)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("нет активной приёмки для ПВЗ %s", pvzID)
		}
		return nil, fmt.Errorf("не удалось закрыть приёмку: %w", err)
	}

	return &reception, nil
}

func (p *pvzRepository) DeleteLastProduct(ctx context.Context, tx pgx.Tx, pvzID string) error {
	query := `
		WITH active_reception AS (
			SELECT reception_id
			FROM receptions
			WHERE pvz_id = $1 AND status = 'in_progress'
			ORDER BY reception_time DESC
			LIMIT 1
		),
		last_product AS (
			SELECT rp.product_id
			FROM reception_products rp
			JOIN products p ON p.product_id = rp.product_id
			JOIN active_reception ar ON ar.reception_id = rp.reception_id
			WHERE p.is_active = TRUE AND rp.is_active = TRUE
			ORDER BY p.reception_time DESC
			LIMIT 1
		),
		update_products AS (
			UPDATE products
			SET is_active = FALSE
			FROM last_product lp
			WHERE products.product_id = lp.product_id
		)
		UPDATE reception_products
		SET is_active = FALSE
		FROM last_product lp, active_reception ar
		WHERE reception_products.product_id = lp.product_id
		  AND reception_products.reception_id = ar.reception_id;
	`

	tag, err := tx.Exec(ctx, query, pvzID)
	if err != nil {
		return fmt.Errorf("ошибка при удалении последнего продукта: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("нет активных продуктов для удаления")
	}
	return nil
}
