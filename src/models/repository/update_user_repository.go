package repository

import (
	"errors"
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
	"focus-finance/src/models/repository/converter"
	"focus-finance/src/models/repository/entity"
	"gorm.io/gorm"
)

func (ur *userRepository) UpdateUserRepository(userDomain models.UserDomainInterface, userID int) *rest_err.RestErr {
	userEntity := converter.ConvertUserDomainToUserEntity(userDomain)
	userEntity.ID = uint(userID)
	err := ur.DB.Model(&userEntity).Updates(entity.User{Name: userEntity.Name, Email: userEntity.Email}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return rest_err.NewInternalServerError("unable update user")

	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return rest_err.NewBadRequestError("user not found")
	}

	return nil
}
