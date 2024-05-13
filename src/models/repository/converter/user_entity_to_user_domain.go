package converter

import (
	"focus-finance/src/models"
	"focus-finance/src/models/repository/entity"
)

func ConvertUserEntityToUserDomain(userEntity entity.User) models.UserDomainInterface {
	userDomain := models.NewUserDomain(userEntity.Name, userEntity.Email, userEntity.Password)
	userDomain.SetID(int(userEntity.ID))

	return userDomain
}
