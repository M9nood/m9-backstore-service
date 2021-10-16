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
	GetProductsInStoreService(storeId *int) ([]product.ProductResponse, iterror.ErrorException)
	CreateProductService(storeId *int, pd product.ProductRequest) (string, iterror.ErrorException)
	UpdateProductService(productUuid string, pd product.ProductRequest) (string, iterror.ErrorException)
	DeleteProductService(productUuid string) (string, iterror.ErrorException)
}

func NewProductService(db *gorm.DB) ProductServiceInterface {
	productRepo := repository.NewProductReposity(db)
	return &ProductService{
		Db:          db,
		ProductRepo: productRepo,
	}
}

func (s ProductService) GetProductsInStoreService(storeId *int) ([]product.ProductResponse, iterror.ErrorException) {
	result, err := s.ProductRepo.GetProducts(storeId)
	if err != nil {
		return product.Products(result).Response(), err
	}
	return product.Products(result).Response(), nil
}

func (s ProductService) CreateProductService(storeId *int, pd product.ProductRequest) (string, iterror.ErrorException) {
	productCreate := pd.ToProductSchema()
	productCreate.StoreId = *storeId
	if _, err := s.ProductRepo.CreateProduct(productCreate); err != nil {
		return "", err
	}
	return "Create product success", nil
}

func (s ProductService) UpdateProductService(productUuid string, pd product.ProductRequest) (string, iterror.ErrorException) {
	productUpdate := pd.ToProductSchema()
	if _, err := s.ProductRepo.UpdateProduct(productUuid, productUpdate); err != nil {
		return "", err
	}
	return "Update product success", nil
}

func (s ProductService) DeleteProductService(productUuid string) (string, iterror.ErrorException) {
	if err := s.ProductRepo.DeleteProduct(productUuid); err != nil {
		return "", err
	}
	return "Delete product success", nil
}
