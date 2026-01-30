package product

import (
	"context"
	"errors"
	"fmt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// ListProducts retrieves a paginated list of products with optional search
func (s *Service) ListProducts(ctx context.Context, params ListParams) (*ProductResponse, error) {
	// Set default limit if not specified
	if params.Limit == 0 {
		params.Limit = 20
	}

	// Clamp values to prevent abuse
	if params.Limit > 100 {
		params.Limit = 100
	}
	if params.Offset < 0 {
		params.Offset = 0
	}

	products, total, err := s.repo.GetProducts(ctx, params.Limit, params.Offset, params.Search)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}

	return &ProductResponse{
		Data:   products,
		Total:  total,
		Limit:  params.Limit,
		Offset: params.Offset,
	}, nil
}

// GetProduct retrieves a single product by ID
func (s *Service) GetProduct(ctx context.Context, pkid int64) (*Product, error) {
	if pkid <= 0 {
		return nil, errors.New("invalid product id")
	}

	product, err := s.repo.GetProductByID(ctx, pkid)
	if err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}

	return product, nil
}
