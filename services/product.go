package service

import (
	"m9-backstore-service/database"

	iterror "github.com/M9nood/go-iterror"
	"github.com/go-pg/pg/v10"
)

type ProductService struct {
	Db *pg.DB
}

var productServiceInstance *ProductService

func NewProductService() *ProductService {
	if productServiceInstance == nil {
		db := database.GetDB()
		productServiceInstance = &ProductService{
			Db: db,
		}
	}
	return productServiceInstance
}

func (pd ProductService) GetProductsService() (string, iterror.ErrorException) {
	return "eieie", nil
}
