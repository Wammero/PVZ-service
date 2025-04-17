package repository

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func (r *Repository) CleanupInactive(ctx context.Context, tx pgx.Tx) error {
	query := `
		WITH deleted_reception_products AS (
			DELETE FROM reception_products
			WHERE is_active = FALSE
			RETURNING 1
		)
		DELETE FROM products
		WHERE is_active = FALSE;
	`

	_, err := r.Pool().Exec(ctx, query)
	return err
}
