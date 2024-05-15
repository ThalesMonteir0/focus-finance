package converter

import (
	"focus-finance/src/models"
	"focus-finance/src/models/repository/entity"
)

func ConvertMovementDomainToEntity(movementDomain models.MovementDomainInterface) entity.Movement {
	return entity.Movement{
		Value:       movementDomain.GetValue(),
		Description: movementDomain.GetDesc(),
		UserID:      movementDomain.GetUserID(),
		TypeID:      movementDomain.GetTypeID(),
		OperationID: movementDomain.GetOperation(),
	}
}
