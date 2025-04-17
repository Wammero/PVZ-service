package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type authRepository struct {
	pool *pgxpool.Pool
}

func NewAuthRepository(pool *pgxpool.Pool) *authRepository {
	return &authRepository{pool: pool}
}

func (r *authRepository) Pool() *pgxpool.Pool {
	return r.pool
}

func (r *authRepository) Register(ctx context.Context, tx pgx.Tx, email, password, salt, userRole string) error {
	query := `
		INSERT INTO users (email, password_hash, salt, role)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (email) DO NOTHING
		RETURNING email;`

	var result string
	err := tx.QueryRow(ctx, query, email, password, salt, userRole).Scan(&result)
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("пользователь с email %s уже зарегистрирован", email)
		}
		return fmt.Errorf("не удалось зарегистрировать пользователя: %v", err)
	}

	return nil
}

func (r *authRepository) Login(ctx context.Context, tx pgx.Tx, email string) (string, string, string, string, error) {
	query := `
		SELECT user_id, password_hash, salt, role
		FROM users
		WHERE email = $1;`

	var id, password, salt, role string
	var err error

	if tx != nil {
		err = tx.QueryRow(ctx, query, email).Scan(&id, &password, &salt, &role)
	} else {
		err = r.pool.QueryRow(ctx, query, email).Scan(&id, &password, &salt, &role)
	}

	if err != nil {
		if err == pgx.ErrNoRows {
			return "", "", "", "", fmt.Errorf("пользователь с email %s не найден", email)
		}
		return "", "", "", "", fmt.Errorf("ошибка при поиске пользователя: %v", err)
	}

	return id, password, salt, role, nil
}
