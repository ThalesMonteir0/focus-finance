package routes

import (
	"focus-finance/src/controller/movements"
	"focus-finance/src/controller/users"
	"focus-finance/src/models"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, userController users.UserControllerInterface, movementController movements.MovementControllerInterface) {
	r.POST("/user", models.VerifyTokenMiddleware, userController.CreateUser)
	r.GET("/user/:id", models.VerifyTokenMiddleware, userController.GetUser)
	r.DELETE("/user/:id", models.VerifyTokenMiddleware, userController.DeleteUser)
	r.PUT("/user/:id", models.VerifyTokenMiddleware, userController.EditUser)
	r.POST("/login", userController.UserLogin)

	r.GET("/movement/:userID", models.VerifyTokenMiddleware, movementController.GetMovement)
	r.POST("/movement", models.VerifyTokenMiddleware, movementController.CreateMovement)
}
