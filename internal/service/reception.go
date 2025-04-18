package service

import (
	"context"
	"fmt"

	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/Wammero/PVZ-service/internal/repository"
)

type receptionService struct {
	repo repository.ReceptionRepository
}

func NewReceptionService(repo repository.ReceptionRepository) *receptionService {
	return &receptionService{repo: repo}
}

func (s *receptionService) CreateReception(ctx context.Context, pvzId string) (*model.Reception, error) {
	tx, err := s.repo.Pool().Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("не удалось начать транзакцию: %v", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			_ = tx.Commit(ctx)
		}
	}()
	reception, err := s.repo.CreateReception(ctx, tx, pvzId)
	if err != nil {
		return nil, err
	}

	return reception, nil
}
