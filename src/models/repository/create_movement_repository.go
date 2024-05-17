package repository

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
	"focus-finance/src/models/repository/converter"
)

func (m *movementRepository) InsertMovement(movementDomain models.MovementDomainInterface) *rest_err.RestErr {
	movementEntity := converter.ConvertMovementDomainToEntity(movementDomain)

	if err := m.DB.Create(&movementEntity).Error; err != nil {
		return rest_err.NewInternalServerError("unable create movement")
	}

	return nil
}
