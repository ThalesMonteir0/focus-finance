package view

import (
	"focus-finance/src/controller/model/response"
	"focus-finance/src/models"
)

func ConvertMovementDomainToResponse(movementDomain models.MovementDomainInterface) response.MovementResponse {
	return response.MovementResponse{
		ID:            movementDomain.GetID(),
		Value:         movementDomain.GetValue(),
		Description:   movementDomain.GetDesc(),
		OperationID:   movementDomain.GetOperation(),
		TypeID:        movementDomain.GetTypeID(),
		UserID:        movementDomain.GetUserID(),
		TypeName:      movementDomain.GetTypeName(),
		OperationName: movementDomain.GetOperationName(),
	}
}
