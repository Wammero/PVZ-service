package service

import (
	"context"

	"github.com/Wammero/PVZ-service/internal/cache"
	"github.com/Wammero/PVZ-service/internal/repository"
)

type pvzService struct {
	repo  repository.PVZRepository
	redis *cache.RedisClient
}

func NewPVZService(repo repository.PVZRepository, redis *cache.RedisClient) *pvzService {
	return &pvzService{repo: repo, redis: redis}
}

func (s *pvzService) CreatePVZ(ctx context.Context, id, registrationDate, city string) error {
	return s.repo.CreatePVZ(ctx, id, registrationDate, city)
}

func (s *pvzService) GetPVZList(ctx context.Context) error {
	return s.repo.GetPVZList(ctx)
}

func (s *pvzService) CloseLastReception(ctx context.Context, pvzID string) error {
	return s.repo.CloseLastReception(ctx, pvzID)
}

func (s *pvzService) DeleteLastProduct(ctx context.Context, pvzID string) error {
	return s.repo.DeleteLastProduct(ctx, pvzID)
}
