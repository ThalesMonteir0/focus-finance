package converter

import (
	"focus-finance/src/models"
	"focus-finance/src/models/repository/entity"
)

func ConvertUserDomainToUserEntity(userDomain models.UserDomainInterface) entity.User {
	return entity.User{
		Email:    userDomain.GetEmail(),
		Password: userDomain.GetPassword(),
		Name:     userDomain.GetName(),
	}
}
