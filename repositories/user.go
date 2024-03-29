package repository

import (
	"m9-backstore-service/models/user"
	util "m9-backstore-service/utils"
	"strings"

	iterror "github.com/M9nood/go-iterror"
	"github.com/jinzhu/gorm"
)

type UserReposity struct {
	Db *gorm.DB
}

type UserReposityInterface interface {
	CreateUser(user user.UserCreateRequest, tx ...*gorm.DB) (user.UserCreateRequest, iterror.ErrorException)
	IsExistUser(username string, email string, tx ...*gorm.DB) (bool, iterror.ErrorException)
	GetDB() *gorm.DB
}

func NewUserReposity(Db *gorm.DB) UserReposityInterface {
	return &UserReposity{
		Db: Db,
	}
}

func (r *UserReposity) GetDB() *gorm.DB {
	return r.Db
}

func (r *UserReposity) CreateUser(user user.UserCreateRequest, tx ...*gorm.DB) (user.UserCreateRequest, iterror.ErrorException) {
	var db *gorm.DB
	if len(tx) > 0 {
		db = tx[0]
	} else {
		db = r.Db
	}
	user.Email = strings.ToLower(user.Email)
	passSult := util.StringRandom(8)
	passHash := util.EncryptSHA1(util.EncryptSHA1(user.Password), passSult)
	user.Password = passHash
	user.PassSault = passSult
	if err := db.Table("user").Create(&user).Error; err != nil {
		return user, iterror.ErrorInternalServer("Error create user")
	}
	return user, nil
}

func (r *UserReposity) IsExistUser(username string, email string, tx ...*gorm.DB) (bool, iterror.ErrorException) {
	var db *gorm.DB
	if len(tx) > 0 {
		db = tx[0]
	} else {
		db = r.Db
	}
	user := []user.UserSchema{}
	if err := db.Table("user").Find(&user, "username = lower(?) OR email = lower(?)", username, email).Error; err != nil {
		return true, iterror.ErrorInternalServer("Error check existing user")
	}
	return len(user) > 0, nil
}
