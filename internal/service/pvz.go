package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Wammero/PVZ-service/internal/cache"
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

func (s *pvzService) GetPVZList(ctx context.Context) error {
	return s.repo.GetPVZList(ctx)
}

func (s *pvzService) CloseLastReception(ctx context.Context, pvzID string) error {
	return s.repo.CloseLastReception(ctx, pvzID)
}

func (s *pvzService) DeleteLastProduct(ctx context.Context, pvzID string) error {
	return s.repo.DeleteLastProduct(ctx, pvzID)
}
