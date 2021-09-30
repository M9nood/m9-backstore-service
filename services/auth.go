package service

import (
	"fmt"
	"m9-backstore-service/models/auth"
	"m9-backstore-service/models/store"
	"m9-backstore-service/models/user"
	repository "m9-backstore-service/repositories"

	iterror "github.com/M9nood/go-iterror"
	"github.com/jinzhu/gorm"
)

type AuthService struct {
	Db        *gorm.DB
	UserRepo  *repository.UserReposity
	StoreRepo *repository.StoreReposity
}

var authServiceInstance *AuthService

func NewAuthService(db *gorm.DB) *AuthService {
	if authServiceInstance == nil {
		userRepo := repository.NewUserReposity(db)
		storeRepo := repository.NewStoreReposity(db)
		authServiceInstance = &AuthService{
			Db:        db,
			UserRepo:  userRepo,
			StoreRepo: storeRepo,
		}
	}
	return authServiceInstance
}

func (s AuthService) RegisterService(register auth.RegisterRequest) (resp string, errSvc iterror.ErrorException) {
	db := s.UserRepo.Db
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	existingUser, err := s.UserRepo.IsExistUser(register.UserName, register.Email, tx)
	if err != nil {
		errSvc = err
		panic(errSvc)
	}
	if existingUser {
		errSvc = iterror.ErrorBadRequest("User was existing")
		panic(errSvc)
	}
	userCreateReq := user.UserCreateRequest{
		UserName: register.UserName,
		Email:    register.Email,
		Password: register.Password,
	}
	resultUser, err := s.UserRepo.CreateUser(userCreateReq, tx)
	if err != nil {
		fmt.Println("error create user", err)
		errSvc = err
		panic(errSvc)
	}
	storeCreate := store.StoreCreateRequest{
		Name: register.StoreName,
		Type: register.StoreType,
	}
	resultStore, err := s.StoreRepo.CreateStore(storeCreate, tx)
	if err != nil {
		fmt.Println("error create store", err)
		errSvc = err
		panic(errSvc)
	}
	fmt.Println("resutltUser", resultUser)
	fmt.Println("resultStore", resultStore)
	tx.Commit()
	return "Register was succesful", nil
}
