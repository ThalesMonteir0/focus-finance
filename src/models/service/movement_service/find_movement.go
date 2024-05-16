package movement_service

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
)

func (m *movementService) GetMovements(movementDomain models.MovementDomainInterface) ([]models.MovementDomainInterface, *rest_err.RestErr) {
	movements, err := m.repository.GetMovements(movementDomain)
	if err != nil {
		return nil, err
	}

	return movements, nil
}
