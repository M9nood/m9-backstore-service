package repository

import (
	"m9-backstore-service/models/product"
	"time"

	iterror "github.com/M9nood/go-iterror"
	"github.com/jinzhu/gorm"
)

type ProductReposity struct {
	Db *gorm.DB
}

var tableName string = "product"

type ProductReposityInterface interface {
	GetProducts(storeId *int) ([]product.ProductSchema, iterror.ErrorException)
	GetProductByCode(code string) (product.ProductSchema, iterror.ErrorException)
	CreateProduct(productData product.ProductSchema) (product.ProductSchema, iterror.ErrorException)
	UpdateProduct(productUuid string, productData product.ProductSchema) (product.ProductSchema, iterror.ErrorException)
	DeleteProduct(productUuid string) iterror.ErrorException
}

func NewProductReposity(Db *gorm.DB) ProductReposityInterface {
	return &ProductReposity{
		Db: Db,
	}
}

func (repo *ProductReposity) GetProducts(storeId *int) ([]product.ProductSchema, iterror.ErrorException) {
	products := []product.ProductSchema{}
	result := repo.Db.Table(tableName).Where("store_id = ?", storeId).Where("delete_flag = 0").Find(&products)
	if result.Error != nil {
		return products, iterror.ErrorInternalServer("Error get products")
	}
	return products, nil
}

func (repo *ProductReposity) GetProductByCode(code string) (product.ProductSchema, iterror.ErrorException) {
	product := product.ProductSchema{}
	if err := repo.Db.Table(tableName).First(&product, "disp_code = ?", code).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return product, iterror.ErrorNotFound("Error not found product")
		}
		return product, iterror.ErrorInternalServer("Error get product")
	}
	return product, nil
}

func (repo *ProductReposity) CreateProduct(productData product.ProductSchema) (product.ProductSchema, iterror.ErrorException) {
	if err := repo.Db.Table(tableName).Create(&productData).Error; err != nil {
		return productData, iterror.ErrorInternalServer("Error create product")
	}
	return productData, nil
}

func (repo *ProductReposity) UpdateProduct(productUuid string, productData product.ProductSchema) (product.ProductSchema, iterror.ErrorException) {
	now := time.Now()
	productData.UpdatedAt = &now
	if err := repo.Db.Table(tableName).Where("product_uuid = ?", productUuid).Update(&productData).Error; err != nil {
		return productData, iterror.ErrorInternalServer("Error update product")
	}
	return productData, nil
}

func (repo *ProductReposity) DeleteProduct(productUuid string) iterror.ErrorException {
	if err := repo.Db.Table(tableName).Where("product_uuid = ?", productUuid).Update(map[string]interface{}{"delete_flag": 1, "updated_at": time.Now()}).Error; err != nil {
		return iterror.ErrorInternalServer("Error delete product")
	}
	return nil
}
