package service

import (
	"m9-backstore-service/models/product"

	repository "m9-backstore-service/repositories"

	iterror "github.com/M9nood/go-iterror"
	"github.com/jinzhu/gorm"
)

type ProductService struct {
	Db          *gorm.DB
	ProductRepo *repository.ProductReposity
}

var productServiceInstance *ProductService

func NewProductService(db *gorm.DB) *ProductService {
	if productServiceInstance == nil {
		productRepo := repository.NewProductReposity(db)
		productServiceInstance = &ProductService{
			Db:          db,
			ProductRepo: productRepo,
		}
	}
	return productServiceInstance
}

func (s ProductService) GetProductsService() ([]product.ProductSchema, iterror.ErrorException) {
	result, err := s.ProductRepo.GetProducts()
	if err != nil {
		return result, err
	}
	return result, nil
}
