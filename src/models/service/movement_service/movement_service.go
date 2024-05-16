package movement_service

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
	"focus-finance/src/models/repository"
)

type movementService struct {
	repository repository.MovementRepositoryInterface
}

type MovementServiceInterface interface {
	GetMovements(models.MovementDomainInterface) ([]models.MovementDomainInterface, *rest_err.RestErr)
	CreateMovement(models.MovementDomainInterface) *rest_err.RestErr
}

func NewMovementService(repository repository.MovementRepositoryInterface) MovementServiceInterface {
	return &movementService{
		repository: repository,
	}
}
