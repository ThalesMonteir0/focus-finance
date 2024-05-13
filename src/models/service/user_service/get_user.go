package user_service

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
)

func (us *userService) GetUserByID(userID int) (models.UserDomainInterface, *rest_err.RestErr) {
	user, err := us.repository.FindUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
