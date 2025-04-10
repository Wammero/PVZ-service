package service

import (
	"context"

	"github.com/Wammero/PVZ-service/internal/cache"
	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/Wammero/PVZ-service/internal/repository"
)

type productService struct {
	repo  repository.ProductRepository
	redis *cache.RedisClient
}

func NewProductService(repo repository.ProductRepository, redis *cache.RedisClient) *productService {
	return &productService{repo: repo, redis: redis}
}

func (s *productService) AddProduct(ctx context.Context, productType model.ProductType, pvzId string) error {
	return s.repo.AddProduct(ctx, productType, pvzId)
}
