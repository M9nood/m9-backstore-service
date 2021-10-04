package repository

import (
	"fmt"
	"m9-backstore-service/models/product"

	iterror "github.com/M9nood/go-iterror"
	"github.com/jinzhu/gorm"
)

type ProductReposity struct {
	Db *gorm.DB
}

type ProductReposityInterface interface {
	GetProducts() ([]product.ProductSchema, iterror.ErrorException)
	GetProductByCode(code string) (product.ProductSchema, iterror.ErrorException)
}

func NewProductReposity(Db *gorm.DB) ProductReposityInterface {
	return &ProductReposity{
		Db: Db,
	}
}

func (repo *ProductReposity) GetProducts() ([]product.ProductSchema, iterror.ErrorException) {
	products := []product.ProductSchema{}
	result := repo.Db.Table("product").Find(&products)
	if result.Error != nil {
		fmt.Println("erro", result.Error)
		return products, iterror.ErrorInternalServer("Error get products")
	}
	return products, nil
}

func (repo *ProductReposity) GetProductByCode(code string) (product.ProductSchema, iterror.ErrorException) {
	product := product.ProductSchema{}
	if err := repo.Db.Table("product").First(&product, "disp_code = ?", code).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return product, iterror.ErrorNotFound("Error not found product")
		}
		return product, iterror.ErrorInternalServer("Error get product")
	}
	return product, nil
}
