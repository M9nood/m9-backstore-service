package service

import (
	"m9-backstore-service/models/product"

	repository "m9-backstore-service/repositories"

	iterror "github.com/M9nood/go-iterror"
	"github.com/jinzhu/gorm"
)

type ProductService struct {
	Db          *gorm.DB
	ProductRepo repository.ProductReposityInterface
}

type ProductServiceInterface interface {
	GetProductsService() ([]product.ProductSchema, iterror.ErrorException)
}

func NewProductService(db *gorm.DB) ProductServiceInterface {
	productRepo := repository.NewProductReposity(db)
	return &ProductService{
		Db:          db,
		ProductRepo: productRepo,
	}
}

func (s ProductService) GetProductsService() ([]product.ProductSchema, iterror.ErrorException) {
	result, err := s.ProductRepo.GetProducts()
	if err != nil {
		return result, err
	}
	return result, nil
}
