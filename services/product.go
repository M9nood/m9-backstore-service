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
	GetProductsInStoreService(storeId *int) ([]product.ProductSchema, iterror.ErrorException)
	CreateProductService(storeId *int, pd product.ProductCreateRequest) (string, iterror.ErrorException)
}

func NewProductService(db *gorm.DB) ProductServiceInterface {
	productRepo := repository.NewProductReposity(db)
	return &ProductService{
		Db:          db,
		ProductRepo: productRepo,
	}
}

func (s ProductService) GetProductsInStoreService(storeId *int) ([]product.ProductSchema, iterror.ErrorException) {
	result, err := s.ProductRepo.GetProducts(storeId)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s ProductService) CreateProductService(storeId *int, pd product.ProductCreateRequest) (string, iterror.ErrorException) {
	productCreate := pd.ToProductSchema()
	productCreate.StoreId = *storeId
	if _, err := s.ProductRepo.CreateProduct(productCreate); err != nil {
		return "", err
	}
	return "Create product success", nil
}
