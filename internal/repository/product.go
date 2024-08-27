package repository

import (
	"example/config"
	"example/entity"
	"example/internal/dto"
)

type (
	ProductRepository interface {
		FindProducts(input *dto.FindProductsInput) ([]*entity.Product, error)
	}

	productRepository struct {
		cfg *config.Config
	}
)

func NewProductRepository(cfg *config.Config) ProductRepository {
	return &productRepository{
		cfg: cfg,
	}
}

func (u *productRepository) FindProducts(input *dto.FindProductsInput) ([]*entity.Product, error) {
	result := []*entity.Product{}
	Product := &entity.Product{
		Name: input.Name,
	}
	result = append(result, Product)
	return result, nil
}
