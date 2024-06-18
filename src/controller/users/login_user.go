package users

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/controller/model/request"
	"focus-finance/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (uc *userController) UserLogin(c *gin.Context) {
	var userlogin request.UserLoginRequest

	if err := c.ShouldBindJSON(&userlogin); err != nil {
		errRequest := rest_err.NewInternalServerError("Unable read json")
		c.JSON(errRequest.Code, errRequest.Message)
		return
	}

	userLoginDomain := models.NewUserLoginDomain(userlogin.Email, userlogin.Password)

	token, err := uc.service.LoginUser(userLoginDomain)
	if err != nil {
		c.JSON(err.Code, err.Message)
		return
	}

	c.JSON(http.StatusOK, token)
}
