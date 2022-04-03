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
	r.GET("/users/:id/applications", controllers.GetUserApplications)
	r.GET("/users/:id/pending", controllers.GetUserPending)
	r.PUT("/users/:id", controllers.UpdateUser)

	r.POST("/applications", controllers.CreateApplication)
	r.GET("/applications/:id", controllers.GetApplication)
	r.PUT("/applications/:id/approve", controllers.ApproveApplication)
	r.PUT("/applications/:id/decline", controllers.DeclineApplication)

	if err := r.Run(); err != nil {
		return
	}
}
