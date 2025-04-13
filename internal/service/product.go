package service

import (
	"context"
	"fmt"

	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/Wammero/PVZ-service/internal/repository"
)

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *productService {
	return &productService{repo: repo}
}

func (s *productService) AddProduct(ctx context.Context, productType string, pvzId string) (*model.Product, error) {
	if !model.IsValidProductType(model.ProductType(productType)) {
		return nil, fmt.Errorf("несуществующая категория продукта")
	}

	tx, err := s.repo.Pool().Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("не удалось начать транзакцию: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			_ = tx.Commit(ctx)
		}
	}()

	return s.repo.AddProduct(ctx, tx, productType, pvzId)
}
