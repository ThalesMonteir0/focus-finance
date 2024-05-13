package users

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/controller/model/request"
	"focus-finance/src/models"
	"focus-finance/src/view"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (uc *userController) CreateUser(c *gin.Context) {
	var user request.UserRequest

	errValidation := c.ShouldBindJSON(&user)
	if errValidation != nil {
		err := rest_err.NewBadRequestError("Erro na validação dos campos")
		c.JSON(err.Code, err.Message)
		return
	}

	userDomain := models.NewUserDomain(user.Name, user.Email, user.Password)
	userCreated, err := uc.service.CreateUser(userDomain)
	if err != nil {
		c.JSON(err.Code, err.Message)
		return
	}

	c.JSON(http.StatusCreated, view.ConvertUserDomainToResponse(userCreated))
	return
}
