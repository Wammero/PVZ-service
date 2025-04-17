package service

import (
	"context"
	"fmt"

	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/Wammero/PVZ-service/internal/repository"
	"github.com/Wammero/PVZ-service/pkg/jwt"
	ps "github.com/Wammero/PVZ-service/pkg/password"
)

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *authService {
	return &authService{repo: repo}
}

func (s *authService) Register(ctx context.Context, email, password string, userRole string) error {
	role := model.UserRole(userRole)
	if !model.IsValidUserRole(role) {
		return fmt.Errorf("invalid role: %s", userRole)
	}

	hashedPassword, salt, err := ps.HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	tx, err := s.repo.Pool().Begin(ctx)
	if err != nil {
		return fmt.Errorf("не удалось начать транзакцию: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			_ = tx.Commit(ctx)
		}
	}()

	err = s.repo.Register(ctx, tx, email, hashedPassword, salt, userRole)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	userID, hashedPassword, salt, role, err := s.repo.Login(ctx, nil, email)
	if err != nil {
		return "", err
	}

	if !ps.CheckPassword(password, salt, hashedPassword) {
		return "", fmt.Errorf("неверный пароль")
	}

	token, err := jwt.GenerateJWT(userID, role)
	if err != nil {
		return "", fmt.Errorf("ошибка генерации токена: %w", err)
	}

	return token, nil
}

func (s *authService) DummyLogin(ctx context.Context, userRole string) (string, error) {
	role := model.UserRole(userRole)
	if !model.IsValidUserRole(role) {
		return "", fmt.Errorf("invalid role: %s", userRole)
	}

	token, err := jwt.GenerateJWT("-1", userRole)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %w", err)
	}

	return token, nil
}
