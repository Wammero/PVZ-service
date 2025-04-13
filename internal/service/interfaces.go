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
	GetPVZList(ctx context.Context, startDateStr, endDateStr, pageStr, limitStr string) ([]model.PVZWithReceptions, error)
	CloseLastReception(ctx context.Context, pvzID string) error
	DeleteLastProduct(ctx context.Context, pvzID string) error
}

type ReceptionService interface {
	CreateReception(ctx context.Context, pvzId string) (string, string, error)
}

type ProductService interface {
	AddProduct(ctx context.Context, productType string, pvzId string) (string, string, string, error)
}
