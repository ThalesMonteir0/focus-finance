package main

import (
	"focus-finance/src/controller/movements"
	"focus-finance/src/controller/users"
	"focus-finance/src/models/repository"
	"focus-finance/src/models/service/movement_service"
	"focus-finance/src/models/service/user_service"
	"gorm.io/gorm"
)

func initDependencies(db *gorm.DB) (users.UserControllerInterface, movements.MovementControllerInterface) {
	userRepo := repository.NewUserRepository(db)
	userService := user_service.NewUserService(userRepo)
	userController := users.NewUserController(userService)

	movementRepo := repository.NewMovementRepository(db)
	movementService := movement_service.NewMovementService(movementRepo)
	movementController := movements.NewMovementController(movementService)

	return userController, movementController
}
