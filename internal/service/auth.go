package service

import (
	"context"

	"github.com/Wammero/PVZ-service/internal/cache"
	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/Wammero/PVZ-service/internal/repository"
)

type authService struct {
	repo  repository.AuthRepository
	redis *cache.RedisClient
}

func NewAuthService(repo repository.AuthRepository, redis *cache.RedisClient) *authService {
	return &authService{repo: repo, redis: redis}
}

func (s *authService) Register(ctx context.Context, email, password string, role model.UserRole) error {
	return s.repo.Register(ctx, email, password, role)
}

func (s *authService) Login(ctx context.Context, email, password string) error {
	return s.repo.Login(ctx, email, password)
}

func (s *authService) DummyLogin(ctx context.Context, userRole model.UserRole) error {
	return nil
}
