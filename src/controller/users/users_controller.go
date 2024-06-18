package users

import (
	"focus-finance/src/models/service/user_service"
	"github.com/gin-gonic/gin"
)

type userController struct {
	service user_service.UserServiceInterface
}

type UserControllerInterface interface {
	CreateUser(*gin.Context)
	DeleteUser(*gin.Context)
	EditUser(*gin.Context)
	GetUser(*gin.Context)
	UserLogin(*gin.Context)
}

func NewUserController(service user_service.UserServiceInterface) UserControllerInterface {
	return &userController{
		service: service,
	}
}
