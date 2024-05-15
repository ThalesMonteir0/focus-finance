package routes

import (
	"focus-finance/src/controller/movements"
	"focus-finance/src/controller/users"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, userController users.UserControllerInterface, movementController movements.MovementControllerInterface) {
	r.POST("/user", userController.CreateUser)
	r.GET("/user/:id", userController.GetUser)
	r.DELETE("/user/:id", userController.DeleteUser)
	r.PUT("/user/:id", userController.EditUser)

	r.GET("/movement", movementController.GetMovement)
	r.POST("/movement", movementController.CreateMovement)
}
