package user_service

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
)

func (us *userService) CreateUser(userDomain models.UserDomainInterface) (models.UserDomainInterface, *rest_err.RestErr) {
	if _, err := us.repository.FindUserByEmail(userDomain.GetEmail()); err == nil {
		err := rest_err.NewBadRequestError("email already exists")
		return nil, err
	}

	userDomain.SetEmailToUppercase()
	userDomain.SetNameToUppercase()
	if err := userDomain.EncryptPassword(); err != nil {
		return nil, err
	}

	userCreated, err := us.repository.CreateUserRepository(userDomain)
	if err != nil {
		return nil, err
	}

	return userCreated, nil
}
