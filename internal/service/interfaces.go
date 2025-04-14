package service

import (
	"context"
	"time"

	"github.com/Wammero/PVZ-service/internal/model"
)

type AuthService interface {
	Register(ctx context.Context, email, password string, userRole model.UserRole) error
	Login(ctx context.Context, email, password string) error
	DummyLogin(ctx context.Context, userRole string) (string, error)
}
type PVZService interface {
	CreatePVZ(ctx context.Context, id, city string, registrationDate time.Time) error
	GetPVZList(ctx context.Context, startDateStr, endDateStr, pageStr, limitStr string) ([]model.PVZWithReceptions, error)
	CloseLastReception(ctx context.Context, pvzID string) (*model.Reception, error)
	DeleteLastProduct(ctx context.Context, pvzID string) error
}

type ReceptionService interface {
	CreateReception(ctx context.Context, pvzId string) (*model.Reception, error)
}

type ProductService interface {
	AddProduct(ctx context.Context, productType string, pvzId string) (*model.Product, error)
}
