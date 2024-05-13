package view

import (
	"focus-finance/src/controller/model/response"
	"focus-finance/src/models"
)

func ConvertUserDomainToResponse(userDomain models.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
	}
}
