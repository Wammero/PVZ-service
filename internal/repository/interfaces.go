package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthRepository interface {
	Register(ctx context.Context, email, password string, userRole model.UserRole) error
	Login(ctx context.Context, email, password string) error
	Pool() *pgxpool.Pool
}

type PVZRepository interface {
	CreatePVZ(ctx context.Context, id, city string, regDate time.Time, creator sql.NullInt64) error
	GetPVZList(ctx context.Context, tx pgx.Tx, startDate, endDate time.Time, page, limit int) ([]model.PVZWithReceptions, error)
	CloseLastReception(ctx context.Context, tx pgx.Tx, pvzID string) (*model.Reception, error)
	DeleteLastProduct(ctx context.Context, tx pgx.Tx, pvzID string) error
	Pool() *pgxpool.Pool
}

type ReceptionRepositor interface {
	CreateReception(ctx context.Context, tx pgx.Tx, pvzId string) (*model.Reception, error)
	Pool() *pgxpool.Pool
}

type ProductRepository interface {
	AddProduct(ctx context.Context, tx pgx.Tx, productType string, pvzId string) (*model.Product, error)
	Pool() *pgxpool.Pool
}
