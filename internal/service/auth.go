package service

import (
	"context"
	"fmt"

	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/Wammero/PVZ-service/internal/repository"
	"github.com/Wammero/PVZ-service/pkg/jwt"
)

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *authService {
	return &authService{repo: repo}
}

func (s *authService) Register(ctx context.Context, email, password string, role model.UserRole) error {
	return s.repo.Register(ctx, email, password, role)
}

func (s *authService) Login(ctx context.Context, email, password string) error {
	return s.repo.Login(ctx, email, password)
}

func (s *authService) DummyLogin(ctx context.Context, userRole string) (string, error) {
	role := model.UserRole(userRole)
	if !model.IsValidUserRole(role) {
		return "", fmt.Errorf("invalid role: %s", userRole)
	}

	token, err := jwt.GenerateJWT(-1, userRole)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %w", err)
	}

	return token, nil
}
