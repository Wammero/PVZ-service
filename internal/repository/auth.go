package repository

import (
	"context"

	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type authRepository struct {
	pool *pgxpool.Pool
}

func NewAuthRepository(pool *pgxpool.Pool) *authRepository {
	return &authRepository{pool: pool}
}

func (a *authRepository) Register(ctx context.Context, email, password string, userRole model.UserRole) error {
	return nil
}

func (a *authRepository) Login(ctx context.Context, email, password string) error {
	return nil
}
