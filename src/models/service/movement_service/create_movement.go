package movement_service

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
)

func (m *movementService) CreateMovement(movementDomain models.MovementDomainInterface) *rest_err.RestErr {
	if err := m.repository.InsertMovement(movementDomain); err != nil {
		return err
	}

	return nil
}
