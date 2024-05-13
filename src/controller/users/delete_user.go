package users

import (
	"focus-finance/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (uc *userController) DeleteUser(c *gin.Context) {
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

	if err := uc.service.DeleteUser(userID); err != nil {
		c.JSON(err.Code, err.Message)
		return
	}

	c.JSON(http.StatusOK, "deleted successful ")
}
