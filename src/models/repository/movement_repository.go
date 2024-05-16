package repository

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
	"gorm.io/gorm"
)

type movementRepository struct {
	DB *gorm.DB
}

type MovementRepositoryInterface interface {
	GetMovements(models.MovementDomainInterface) ([]models.MovementDomainInterface, *rest_err.RestErr)
	InsertMovement(models.MovementDomainInterface) *rest_err.RestErr
}

func NewMovementRepository(DB *gorm.DB) MovementRepositoryInterface {
	return &movementRepository{
		DB: DB,
	}
}
