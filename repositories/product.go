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

func NewProductReposity(Db *gorm.DB) *ProductReposity {
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
	// if err := repo.Db.Table("product").Find(&products).Error; err != nil {
	// 	return products, iterror.ErrorInternalServer("Error get products")
	// }
	return products, nil
}
