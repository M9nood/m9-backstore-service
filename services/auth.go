package service

import (
	"m9-backstore-service/models/auth"
	repository "m9-backstore-service/repositories"

	iterror "github.com/M9nood/go-iterror"
	"github.com/jinzhu/gorm"
)

type AuthService struct {
	Db       *gorm.DB
	UserRepo *repository.UserReposity
}

var authServiceInstance *AuthService

func NewAuthService(db *gorm.DB) *AuthService {
	if authServiceInstance == nil {
		userRepo := repository.NewUserReposity(db)
		authServiceInstance = &AuthService{
			Db:       db,
			UserRepo: userRepo,
		}
	}
	return authServiceInstance
}

func (s AuthService) RegisterService(register auth.RegisterRequest) (string, iterror.ErrorException) {
	// result, err := s.ProductRepo.GetProducts()
	// if err != nil {
	// 	return result, err
	// }
	// return result, nil
	return "Register was succesful", nil
}
