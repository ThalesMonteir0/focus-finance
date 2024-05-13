package user_service

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
	"focus-finance/src/models/repository"
)

type userService struct {
	repository repository.UserRepositoryInterface
}

type UserServiceInterface interface {
	CreateUser(models.UserDomainInterface) (models.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(models.UserDomainInterface) *rest_err.RestErr
	DeleteUser(int) *rest_err.RestErr
	GetUserByID(int) (models.UserDomainInterface, *rest_err.RestErr)
}

func NewUserService(repository repository.UserRepositoryInterface) UserServiceInterface {
	return &userService{
		repository: repository,
	}
}
