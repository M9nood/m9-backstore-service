package service

import (
	"fmt"
	"m9-backstore-service/models/auth"
	"m9-backstore-service/models/store"
	"m9-backstore-service/models/user"
	repository "m9-backstore-service/repositories"
	util "m9-backstore-service/utils"

	iterror "github.com/M9nood/go-iterror"
	"github.com/jinzhu/gorm"
)

type AuthService struct {
	Db        *gorm.DB
	UserRepo  repository.UserReposityInterface
	StoreRepo repository.StoreReposityInterface
}

type AuthServiceInterface interface {
	RegisterService(register auth.RegisterRequest) (resp string, errSvc iterror.ErrorException)
	LoginService(user auth.LoginRequest) (auth.LoginResponse, iterror.ErrorException)
	RefreshTokenService(userId int) (auth.Token, iterror.ErrorException)
}

func NewAuthService(db *gorm.DB) AuthServiceInterface {
	userRepo := repository.NewUserReposity(db)
	storeRepo := repository.NewStoreReposity(db)
	return &AuthService{
		Db:        db,
		UserRepo:  userRepo,
		StoreRepo: storeRepo,
	}
}

func (s AuthService) RegisterService(register auth.RegisterRequest) (resp string, errSvc iterror.ErrorException) {
	db := s.Db
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
	storeOwnerCreate := store.StoreOwnerCreateRequest{
		OwnerType: 1,
		UserId:    resultUser.Id,
		StoreId:   resultStore.Id,
	}
	if _, err := s.StoreRepo.CreateStoreOwner(storeOwnerCreate, tx); err != nil {
		fmt.Println("error create store owner", err)
		errSvc = err
		panic(errSvc)
	}
	tx.Commit()
	return "Register was succesful", nil
}

func (s AuthService) LoginService(user auth.LoginRequest) (auth.LoginResponse, iterror.ErrorException) {
	resp := auth.LoginResponse{}
	userFound, err := s.UserRepo.FindByUsernameAndEmail(user.UserName, user.UserName)
	if err != nil {
		return resp, err
	}
	dbPassword := userFound.Password
	passHash := util.EncryptSHA1(user.PasswordHash, userFound.PassSault)
	if passHash != dbPassword {
		return resp, iterror.ErrorBadRequest("Invalid username or password")
	}
	storeId, err := s.StoreRepo.FindStoreIdFromUser(userFound.Id)
	if err != nil {
		return resp, err
	}
	payload := auth.LoginResponse{
		Id:       userFound.Id,
		UserName: userFound.UserName,
		Email:    userFound.Email,
		StoreId:  storeId,
	}

	jwtSvc := NewJWTAuthService()
	accessToken := jwtSvc.GenerateToken(payload)
	refreshToken := jwtSvc.GenerateRefreshToken(userFound.Id)

	resp = payload
	resp.Token = auth.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return resp, nil
}

func (s AuthService) RefreshTokenService(userId int) (auth.Token, iterror.ErrorException) {
	resp := auth.Token{}
	userFound, err := s.UserRepo.FindById(userId)
	if err != nil {
		return resp, err
	}
	storeId, err := s.StoreRepo.FindStoreIdFromUser(userId)
	if err != nil {
		return resp, err
	}
	payload := auth.LoginResponse{
		Id:       userId,
		UserName: userFound.UserName,
		Email:    userFound.Email,
		StoreId:  storeId,
	}

	jwtSvc := NewJWTAuthService()
	accessToken := jwtSvc.GenerateToken(payload)
	refreshToken := jwtSvc.GenerateRefreshToken(userId)
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken
	return resp, nil
}
