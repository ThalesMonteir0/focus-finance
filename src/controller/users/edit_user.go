package users

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/controller/model/request"
	"focus-finance/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (uc *userController) EditUser(c *gin.Context) {
	var userRequest request.UserUpdateRequest
	userID, errConvert := strconv.Atoi(c.Param("id"))
	if errConvert != nil {
		err := rest_err.NewInternalServerError("unable get param")
		c.JSON(err.Code, err.Message)
		return
	}

	if userID == 0 {
		err := rest_err.NewBadRequestError("user ID not null")
		c.JSON(err.Code, err.Message)
		return
	}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRequest := rest_err.NewBadRequestError("param not found")
		c.JSON(errRequest.Code, errRequest.Message)
	}

	userDomain := models.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Email)

	if err := uc.service.UpdateUser(userDomain, userID); err != nil {
		c.JSON(err.Code, err.Message)
		return
	}

	c.JSON(http.StatusOK, "user updated successful")
	return
}
