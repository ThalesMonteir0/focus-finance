package main

import (
	"focus-finance/src/configuration/database"
	"focus-finance/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {

	}

	DB := database.NewPostgresDB()
	userController, movementController := initDependencies(DB)

	router := gin.Default()
	routes.Routes(router, userController, movementController)

	if err := router.Run(":5000"); err != nil {

	}
}
