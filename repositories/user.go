package repository

import (
	"m9-backstore-service/models/user"

	iterror "github.com/M9nood/go-iterror"
	"github.com/jinzhu/gorm"
)

type UserReposity struct {
	Db *gorm.DB
}

func NewUserReposity(Db *gorm.DB) *UserReposity {
	return &UserReposity{
		Db: Db,
	}
}

func (r *UserReposity) CreateUser(user user.UserSchema, tx ...*gorm.DB) (user.UserSchema, iterror.ErrorException) {
	var db *gorm.DB
	if len(tx) > 0 {
		db = tx[0]
	} else {
		db = r.Db
	}
	if err := db.Table("user").Create(&user).Error; err != nil {
		return user, iterror.ErrorInternalServer("Error create user")
	}
	return user, nil
}
