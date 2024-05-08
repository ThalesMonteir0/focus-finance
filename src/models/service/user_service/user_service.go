package user_service

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
)

type userService struct {
	repository string
}

type UserServiceInterface interface {
	CreateUser(models.UserDomainInterface) (models.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(models.UserDomainInterface) *rest_err.RestErr
	DeleteUser(int) *rest_err.RestErr
	GetUserByID(int) (models.UserDomainInterface, *rest_err.RestErr)
}

func NewUserService(repository string) UserServiceInterface {
	return &userService{
		repository: repository,
	}
}
