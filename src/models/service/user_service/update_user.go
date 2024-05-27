package user_service

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
)

func (us *userService) UpdateUser(userDomain models.UserDomainInterface, userID int) *rest_err.RestErr {
	userDomain.SetNameToUppercase()
	userDomain.SetEmailToUppercase()
	if err := userDomain.EncryptPassword(); err != nil {
		return rest_err.NewInternalServerError("unable encrypt password")
	}

	if err := us.repository.UpdateUserRepository(userDomain, userID); err != nil {
		return err
	}

	return nil
}
