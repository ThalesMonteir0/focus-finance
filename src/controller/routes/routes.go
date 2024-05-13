package routes

import (
	"focus-finance/src/controller/users"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, userController users.UserControllerInterface) {
	r.GET("/version")
	r.POST("/user", userController.CreateUser)
}
