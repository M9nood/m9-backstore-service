package repository

import (
	"m9-backstore-service/models/store"

	iterror "github.com/M9nood/go-iterror"
	"github.com/jinzhu/gorm"
)

type StoreReposity struct {
	Db *gorm.DB
}

func NewStoreReposity(Db *gorm.DB) *StoreReposity {
	return &StoreReposity{
		Db: Db,
	}
}

func (r *StoreReposity) CreateStore(store store.StoreCreateRequest, tx ...*gorm.DB) (store.StoreCreateRequest, iterror.ErrorException) {
	var db *gorm.DB
	if len(tx) > 0 {
		db = tx[0]
	} else {
		db = r.Db
	}
	if err := db.Table("store").Create(&store).Error; err != nil {
		return store, iterror.ErrorInternalServer("Error create store")
	}
	return store, nil
}
