package service

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/Wammero/PVZ-service/internal/cache"
	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/Wammero/PVZ-service/internal/repository"
	"github.com/Wammero/PVZ-service/pkg/jwt"
)

type pvzService struct {
	repo  repository.PVZRepository
	redis *cache.RedisClient
}

func NewPVZService(repo repository.PVZRepository, redis *cache.RedisClient) *pvzService {
	return &pvzService{repo: repo, redis: redis}
}

func (s *pvzService) CreatePVZ(ctx context.Context, id, registrationDate, city string) error {
	switch city {
	case "Moscow", "Saint-Petersburg", "Kazan":
	default:
		return fmt.Errorf("недопустимый город: %s", city)
	}

	creatorID, ok := jwt.GetUserID(ctx)
	if !ok {
		return fmt.Errorf("не удалось получить userID из JWT токена")
	}

	var creator sql.NullInt64
	if creatorID != -1 {
		creator = sql.NullInt64{Int64: int64(creatorID), Valid: true}
	} else {
		creator = sql.NullInt64{Valid: false}
	}

	regDate, err := time.Parse(time.RFC3339, registrationDate)
	if err != nil {
		return fmt.Errorf("не удалось парсить дату регистрации: %v", err)
	}

	err = s.repo.CreatePVZ(ctx, id, city, regDate, creator)

	return err
}

func (s *pvzService) GetPVZList(ctx context.Context, startDateStr, endDateStr, pageStr, limitStr string) ([]model.PVZWithReceptions, error) {
	startDate, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid startDate format: %w", err)
	}

	endDate, err := time.Parse(time.RFC3339, endDateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid endDate format: %w", err)
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

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

	pvzList, err := s.repo.GetPVZList(ctx, tx, startDate, endDate, page, limit)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить данные о pvz: %v", err)
	}

	return pvzList, nil
}

func (s *pvzService) CloseLastReception(ctx context.Context, pvzID string) error {
	tx, err := s.repo.Pool().Begin(ctx)
	if err != nil {
		return fmt.Errorf("не удалось начать транзакцию: %v", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			_ = tx.Commit(ctx)
		}
	}()

	err = s.repo.CloseLastReception(ctx, tx, pvzID)
	return err
}

func (s *pvzService) DeleteLastProduct(ctx context.Context, pvzID string) error {
	tx, err := s.repo.Pool().Begin(ctx)
	if err != nil {
		return fmt.Errorf("не удалось начать транзакцию: %v", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			_ = tx.Commit(ctx)
		}
	}()

	err = s.repo.DeleteLastProduct(ctx, tx, pvzID)

	return err
}
