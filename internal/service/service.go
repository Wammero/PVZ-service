package service

import (
	"github.com/Wammero/PVZ-service/internal/repository"
)

type Service struct {
	AuthService      AuthService
	PVZService       PVZService
	ReceptionService ReceptionService
	ProductService   ProductService
}

func New(repo *repository.Repository) *Service {
	return &Service{
		AuthService:      NewAuthService(repo.AuthRepository),
		PVZService:       NewPVZService(repo.PVZRepository),
		ReceptionService: NewReceptionService(repo.ReceptionRepositor),
		ProductService:   NewProductService(repo.ProductRepository),
	}
}
