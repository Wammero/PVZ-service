package service

import (
	"github.com/Wammero/PVZ-service/internal/cache"
	"github.com/Wammero/PVZ-service/internal/repository"
)

type Service struct {
	repo  *repository.Repository
	redis *cache.RedisClient
}

func New(repo *repository.Repository, redisClient *cache.RedisClient) *Service {
	return &Service{
		repo:  repo,
		redis: redisClient,
	}
}
