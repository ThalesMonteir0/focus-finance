package movements

import (
	"focus-finance/src/controller/model/request"
	"focus-finance/src/controller/model/response"
	"focus-finance/src/models"
	"focus-finance/src/view"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (m *movementController) GetMovement(c *gin.Context) {
	movementRequest := getParams(c)

	movementDomain := models.NewMovementsDomain(movementRequest.Description, movementRequest.Value,
		movementRequest.OperationID,
		movementRequest.TypeID,
		movementRequest.UserID)

	movements, err := m.service.GetMovements(movementDomain)
	if err != nil {
		c.JSON(err.Code, err.Message)
		return
	}

	movementsResponseList := formatMovementDomainListToMovementResponseList(movements)

	c.JSON(http.StatusOK, movementsResponseList)
	return
}

func getParams(c *gin.Context) request.MovementRequest {
	var movementRequest request.MovementRequest

	movementRequest.TypeID, _ = strconv.Atoi(c.Query("type_id"))
	movementRequest.OperationID, _ = strconv.Atoi(c.Query("operation_id"))
	movementRequest.UserID, _ = strconv.Atoi(c.Param("userID"))

	return movementRequest
}

func formatMovementDomainListToMovementResponseList(movements []models.MovementDomainInterface) []response.MovementResponse {
	var movementsResponse []response.MovementResponse
	for _, movement := range movements {
		movementsResponse = append(movementsResponse, view.ConvertMovementDomainToResponse(movement))
	}

	return movementsResponse
}
