package repository

import (
	"context"

	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthRepository struct {
	pool *pgxpool.Pool
}

func NewAuthRepository(pool *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{pool: pool}
}

func (a *AuthRepository) DummyLogin(ctx context.Context, userRole model.UserRole) error {
	return nil
}

func (a *AuthRepository) Register(ctx context.Context, email, password string, userRole model.UserRole) error {
	return nil
}

func (a *AuthRepository) Login(ctx context.Context, email, password string) error {
	return nil
}
