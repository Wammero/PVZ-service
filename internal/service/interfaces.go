package service

import (
	"context"

	"github.com/Wammero/PVZ-service/internal/model"
)

type AuthService interface {
	Register(ctx context.Context, email, password string, userRole model.UserRole) error
	Login(ctx context.Context, email, password string) error
	DummyLogin(ctx context.Context, userRole string) (string, error)
}
type PVZService interface {
	CreatePVZ(ctx context.Context, id, registrationDate, city string) error
	GetPVZList(ctx context.Context) error
	CloseLastReception(ctx context.Context, pvzID string) error
	DeleteLastProduct(ctx context.Context, pvzID string) error
}

type ReceptionService interface {
	CreateReception(ctx context.Context, pvzId string) error
}

type ProductService interface {
	AddProduct(ctx context.Context, productType model.ProductType, pvzId string) error
}
