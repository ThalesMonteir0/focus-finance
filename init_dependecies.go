package main

import (
	"focus-finance/src/controller/users"
	"focus-finance/src/models/repository"
	"focus-finance/src/models/service/user_service"
	"gorm.io/gorm"
)

func initDependencies(db *gorm.DB) users.UserControllerInterface {
	userRepo := repository.NewUserRepository(db)
	userService := user_service.NewUserService(userRepo)
	userController := users.NewUserController(userService)

	return userController
}
