package repository

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepositoryInterface interface {
	CreateUserRepository(models.UserDomainInterface) (models.UserDomainInterface, *rest_err.RestErr)
	UpdateUserRepository(models.UserDomainInterface) (int, *rest_err.RestErr)
	DeleteUserRepository(int) *rest_err.RestErr
	FindUserByID(int) (models.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(string) (models.UserDomainInterface, *rest_err.RestErr)
}

func NewUserRepository(DB *gorm.DB) UserRepositoryInterface {
	return &userRepository{
		DB: DB,
	}
}
