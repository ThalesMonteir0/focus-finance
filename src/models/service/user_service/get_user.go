package user_service

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
)

func (us *userService) GetUserByID(int) (models.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}