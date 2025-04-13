package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Wammero/PVZ-service/internal/model"
)

type AuthRepository interface {
	Register(ctx context.Context, email, password string, userRole model.UserRole) error
	Login(ctx context.Context, email, password string) error
}

type PVZRepository interface {
	CreatePVZ(ctx context.Context, id, city string, regDate time.Time, creator sql.NullInt64) error
	GetPVZList(ctx context.Context) error
	CloseLastReception(ctx context.Context, pvzID string) error
	DeleteLastProduct(ctx context.Context, pvzID string) error
}

type ReceptionRepositor interface {
	CreateReception(ctx context.Context, pvzId string) error
}

type ProductRepository interface {
	AddProduct(ctx context.Context, productType model.ProductType, pvzId string) error
}
