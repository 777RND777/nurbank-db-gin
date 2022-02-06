package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nur-bank-gin-api/models"
)

func CreateApplication(c *gin.Context) {
	var input models.CreateApplicationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Application{})
		return
	}

	var application = models.Application{
		ID:          input.ID,
		UserID:      input.UserID,
		Value:       input.Value,
		RequestDate: input.RequestDate,
		IsAdmin:     input.IsAdmin,
	}

	models.DB.Create(application)

	c.JSON(http.StatusOK, application)
}

func GetApplicationList(c *gin.Context) {
	var applications []models.Application
	models.DB.Find(&applications)

	c.JSON(http.StatusOK, applications)
}

func GetApplication(c *gin.Context) {
	var application models.Application
	if err := models.DB.Where("id = ?", c.Param("id")).First(&application).Error; err != nil {
		application.PK = -1
		c.JSON(http.StatusBadRequest, application)
		return
	}

	c.JSON(http.StatusOK, application)
}

func UpdateApplication(c *gin.Context) {
	var application models.Application
	if err := models.DB.Where("id = ?", c.Param("id")).First(&application).Error; err != nil {
		c.JSON(http.StatusBadRequest, application)
		return
	}

	var input models.UpdateApplicationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		//never goes here
		c.JSON(http.StatusBadRequest, application)
		return
	}

	models.DB.Model(&application).Update(input)

	c.JSON(http.StatusOK, application)
}

func RemoveApplication(c *gin.Context) {
	var application models.Application
	if err := models.DB.Where("id = ?", c.Param("id")).First(&application).Error; err != nil {
		c.JSON(http.StatusBadRequest, application)
		return
	}

	models.DB.Delete(&application)

	c.JSON(http.StatusOK, true)
}
