package converter

import (
	"focus-finance/src/models"
	"focus-finance/src/models/repository/entity"
)

func ConvertMovementEntityToDomain(movement entity.Movement) models.MovementDomainInterface {
	movementDomain := models.NewMovementsDomain(
		movement.Description, movement.Value,
		movement.OperationID,
		movement.TypeID,
		movement.UserID)

	movementDomain.SetID(int(movement.ID))
	movementDomain.SetOperationName(movement.Operation.Name)
	movementDomain.SetTypeName(movement.Type.Name)

	return movementDomain
}
