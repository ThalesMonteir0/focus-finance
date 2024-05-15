package movements

import (
	"focus-finance/src/models/service/movement_service"
	"github.com/gin-gonic/gin"
)

type movementController struct {
	service movement_service.MovementServiceInterface
}

type MovementControllerInterface interface {
	GetMovement(*gin.Context)
	CreateMovement(*gin.Context)
}

func NewMovementController(service movement_service.MovementServiceInterface) MovementControllerInterface {
	return &movementController{
		service: service,
	}
}
