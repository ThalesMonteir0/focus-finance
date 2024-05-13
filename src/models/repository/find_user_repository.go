package repository

import (
	"errors"
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
	"focus-finance/src/models/repository/converter"
	"focus-finance/src/models/repository/entity"
	"gorm.io/gorm"
)

func (ur *userRepository) FindUserByID(id int) (models.UserDomainInterface, *rest_err.RestErr) {
	var user entity.User

	err := ur.DB.First(&user, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		err := rest_err.NewInternalServerError("unable get user")
		return nil, err
	}

	if user.ID == 0 || errors.Is(err, gorm.ErrRecordNotFound) {
		err := rest_err.NewInternalServerError("user not found")
		return nil, err
	}

	return converter.ConvertUserEntityToUserDomain(user), nil
}

func (ur *userRepository) FindUserByEmail(email string) (models.UserDomainInterface, *rest_err.RestErr) {
	var user entity.User

	err := ur.DB.Where("email = ?", email).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		err := rest_err.NewInternalServerError("unable get user")
		return nil, err
	}

	if user.ID == 0 || errors.Is(err, gorm.ErrRecordNotFound) {
		err := rest_err.NewBadRequestError("user not found")
		return nil, err
	}

	return converter.ConvertUserEntityToUserDomain(user), nil
}
