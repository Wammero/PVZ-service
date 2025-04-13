package service

import (
	"context"
	"fmt"

	"github.com/Wammero/PVZ-service/internal/cache"
	"github.com/Wammero/PVZ-service/internal/repository"
)

type receptionService struct {
	repo  repository.ReceptionRepositor
	redis cache.RedisClient
}

func NewReceptionService(repo repository.ReceptionRepositor) *receptionService {
	return &receptionService{repo: repo}
}

func (s *receptionService) CreateReception(ctx context.Context, pvzId string) (string, string, error) {
	tx, err := s.repo.Pool().Begin(ctx)
	if err != nil {
		return "", "", fmt.Errorf("не удалось начать транзакцию: %v", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			_ = tx.Commit(ctx)
		}
	}()
	receptionId, dateTime, err := s.repo.CreateReception(ctx, tx, pvzId)
	if err != nil {
		return "", "", err
	}

	return receptionId, dateTime, nil
}
