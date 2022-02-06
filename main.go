package main

import (
	"github.com/gin-gonic/gin"

	"nur-bank-gin-api/controllers"
	"nur-bank-gin-api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	r.GET("/users", controllers.GetUsers)

	if err := r.Run(); err != nil {
		return
	}
}
