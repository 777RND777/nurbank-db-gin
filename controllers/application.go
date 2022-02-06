package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nur-bank-gin-api/models"
)

func CreateApplication(c *gin.Context) {
	var input models.CreateApplicationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	c.JSON(http.StatusOK, gin.H{"data": application})
}

func GetApplicationList(c *gin.Context) {
	var applications []models.Application
	models.DB.Find(&applications)

	c.JSON(http.StatusOK, gin.H{"data": applications})
}

func GetApplication(c *gin.Context) {
	var application models.Application
	if err := models.DB.Where("id = ?", c.Param("id")).First(&application).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": application})
}

func UpdateApplication(c *gin.Context) {
	var application models.Application
	if err := models.DB.Where("id = ?", c.Param("id")).First(&application).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdateApplicationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		//never goes here
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&application).Update(input)

	c.JSON(http.StatusOK, gin.H{"data": application})
}

func RemoveApplication(c *gin.Context) {
	var application models.Application
	if err := models.DB.Where("id = ?", c.Param("id")).First(&application).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&application)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
