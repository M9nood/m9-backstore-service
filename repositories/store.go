package repository

import (
	"m9-backstore-service/models/store"

	iterror "github.com/M9nood/go-iterror"
	"github.com/jinzhu/gorm"
)

type StoreReposity struct {
	Db *gorm.DB
}

type StoreReposityInterface interface {
	CreateStore(store store.StoreCreateRequest, tx ...*gorm.DB) (store.StoreCreateRequest, iterror.ErrorException)
	CreateStoreOwner(owner store.StoreOwnerCreateRequest, tx ...*gorm.DB) (store.StoreOwnerCreateRequest, iterror.ErrorException)
}

func NewStoreReposity(Db *gorm.DB) StoreReposityInterface {
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

func (r *StoreReposity) CreateStoreOwner(owner store.StoreOwnerCreateRequest, tx ...*gorm.DB) (store.StoreOwnerCreateRequest, iterror.ErrorException) {
	var db *gorm.DB
	if len(tx) > 0 {
		db = tx[0]
	} else {
		db = r.Db
	}
	if err := db.Table("store_owner").Create(&owner).Error; err != nil {
		return owner, iterror.ErrorInternalServer("Error create store owner")
	}
	return owner, nil
}
