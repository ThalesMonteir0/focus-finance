package main

import (
	"focus-finance/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {

	}

	router := gin.Default()
	routes.Routes(router)

	if err := router.Run(":5000"); err != nil {

	}
}
