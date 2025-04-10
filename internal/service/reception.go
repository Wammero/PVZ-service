package service

import (
	"context"

	"github.com/Wammero/PVZ-service/internal/cache"
	"github.com/Wammero/PVZ-service/internal/repository"
)

type receptionService struct {
	repo  repository.ReceptionRepositor
	redis *cache.RedisClient
}

func NewReceptionService(repo repository.ReceptionRepositor, redis *cache.RedisClient) *receptionService {
	return &receptionService{repo: repo, redis: redis}
}

func (s *receptionService) CreateReception(ctx context.Context, pvzId string) error {
	return s.repo.CreateReception(ctx, pvzId)
}
