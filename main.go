package main

import (
	"github.com/gin-gonic/gin"

	"nur-bank-gin-api/controllers"
	"nur-bank-gin-api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUserList)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.RemoveUser)

	r.POST("/applications", controllers.CreateApplication)
	r.GET("/applications", controllers.GetApplicationList)
	r.GET("/applications/:id", controllers.GetApplication)
	r.PUT("/applications/:id", controllers.UpdateApplication)
	r.DELETE("/applications/:id", controllers.RemoveApplication)

	if err := r.Run(); err != nil {
		return
	}
}
