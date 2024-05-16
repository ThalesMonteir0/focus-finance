package repository

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
	"focus-finance/src/models/repository/converter"
	"focus-finance/src/models/repository/entity"
)

func (m *movementRepository) GetMovements(movementDomain models.MovementDomainInterface) ([]models.MovementDomainInterface, *rest_err.RestErr) {
	var result []entity.Movement
	var movementListDomain []models.MovementDomainInterface
	movementEntity := converter.ConvertMovementDomainToEntity(movementDomain)

	err := m.DB.Preload("Type").Preload("Operation").
		Where(&movementEntity).
		Find(&result).Error
	if err != nil {
		return nil, rest_err.NewInternalServerError("unable get movement")
	}

	for _, movement := range result {
		movementListDomain = append(movementListDomain, converter.ConvertMovementEntityToDomain(movement))
	}

	return movementListDomain, nil
}
