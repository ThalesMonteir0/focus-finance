package repository

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
	"focus-finance/src/models/repository/converter"
)

func (ur *userRepository) CreateUserRepository(userDomain models.UserDomainInterface) (models.UserDomainInterface, *rest_err.RestErr) {
	userEntity := converter.ConvertUserDomainToUserEntity(userDomain)

	result := ur.DB.Create(&userEntity)
	if result.RowsAffected == 0 {
		err := rest_err.NewInternalServerError("Unable to create user")
		return nil, err
	}

	return converter.ConvertUserEntityToUserDomain(userEntity), nil
}
