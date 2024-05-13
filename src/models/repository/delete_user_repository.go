package repository

import (
	"errors"
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models/repository/entity"
	"gorm.io/gorm"
)

func (ur *userRepository) DeleteUserRepository(id int) *rest_err.RestErr {
	var user entity.User

	err := ur.DB.Delete(&user, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		err := rest_err.NewInternalServerError("unable delete user")
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err := rest_err.NewBadRequestError("user not found")
		return err
	}

	return nil
}
