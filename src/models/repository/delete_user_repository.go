package repository

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models/repository/entity"
)

func (ur *userRepository) DeleteUserRepository(id int) *rest_err.RestErr {
	var user entity.User

	err := ur.DB.Delete(&user, id).Error
	if err != nil {
		//todo:logs
		err := rest_err.NewInternalServerError("unable delete user")
		return err
	}

	return nil
}
