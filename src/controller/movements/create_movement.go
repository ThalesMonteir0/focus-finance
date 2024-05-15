package movements

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/controller/model/request"
	"focus-finance/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (m *movementController) CreateMovement(c *gin.Context) {
	var movementRequest request.MovementRequest

	if err := c.ShouldBindJSON(&movementRequest); err != nil {
		badRequestERR := rest_err.NewBadRequestError("all fields are mandatory")
		c.JSON(badRequestERR.Code, badRequestERR.Message)
		return
	}

	movementDomain := models.NewMovementsDomain(
		movementRequest.Description,
		movementRequest.Value,
		movementRequest.OperationID,
		movementRequest.TypeID,
		movementRequest.UserID)

	if err := m.service.CreateMovement(movementDomain); err != nil {
		c.JSON(err.Code, err.Message)
		return
	}

	c.JSON(http.StatusOK, "movement successfully created")
}
