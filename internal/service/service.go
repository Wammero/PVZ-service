package service

import (
	"github.com/Wammero/PVZ-service/internal/cache"
	"github.com/Wammero/PVZ-service/internal/repository"
)

type Service struct {
	AuthService      AuthService
	PVZService       PVZService
	ReceptionService ReceptionService
	ProductService   ProductService
}

func New(repo *repository.Repository, redisClient *cache.RedisClient) *Service {
	return &Service{
		AuthService:      NewAuthService(repo.AuthRepository, redisClient),
		PVZService:       NewPVZService(repo.PVZRepository, redisClient),
		ReceptionService: NewReceptionService(repo.ReceptionRepositor, redisClient),
		ProductService:   NewProductService(repo.ProductRepository, redisClient),
	}
}
