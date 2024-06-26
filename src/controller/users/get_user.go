package users

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/view"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (uc *userController) GetUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	if userID == 0 {
		err := rest_err.NewBadRequestError("user ID not null")
		c.JSON(err.Code, err.Message)
		return
	}

	user, err := uc.service.GetUserByID(userID)
	if err != nil {
		c.JSON(err.Code, err.Message)
		return
	}

	c.JSON(http.StatusOK, view.ConvertUserDomainToResponse(user))
	return
}
