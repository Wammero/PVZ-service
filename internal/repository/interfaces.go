package repository

import (
	"context"

	"github.com/Wammero/PVZ-service/internal/model"
)

type Authorization interface {
	DummyLogin(ctx context.Context, userRole model.UserRole) error
	Register(ctx context.Context, email, password string, userRole model.UserRole) error
	Login(ctx context.Context, email, password string) error
}

type PVZ interface {
	CreatePVZ(ctx context.Context, id, registrationDate, city string) error
	GetPVZList(ctx context.Context) error
	CloseLastReception(ctx context.Context, pvzID string) error
	DeleteLastProduct(ctx context.Context, pvzID string) error
}

type Reception interface {
	CreateReception(ctx context.Context, pvzId string) error
}

type Product interface {
	AddProduct(ctx context.Context, productType model.ProductType, pvzId string) error
}
