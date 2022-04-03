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
		Password:    input.Password,
		Value:       input.Value,
		RequestDate: GetCurrentTime(),
		IsAdmin:     input.IsAdmin,
	}

	models.DB.Create(application)

	c.JSON(http.StatusOK, application)
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

func ApproveApplication(c *gin.Context) {
	var application models.Application
	if err := models.DB.Where("id = ?", c.Param("id")).First(&application).Error; err != nil {
		application.PK = -1
		c.JSON(http.StatusBadRequest, application)
		return
	}
	if len(application.AnswerDate) > 0 {
		c.JSON(http.StatusBadRequest, application)
		return
	}

	application.AnswerDate = GetCurrentTime()
	application.Approved = true
	models.DB.Model(&application).Update()

	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, user)
		return
	}

	user.Debt += application.Value
	models.DB.Save(&user)

	c.JSON(http.StatusOK, application)
}

func DeclineApplication(c *gin.Context) {
	var application models.Application
	if err := models.DB.Where("id = ?", c.Param("id")).First(&application).Error; err != nil {
		application.PK = -1
		c.JSON(http.StatusBadRequest, application)
		return
	}
	if len(application.AnswerDate) > 0 {
		c.JSON(http.StatusBadRequest, application)
		return
	}

	application.AnswerDate = GetCurrentTime()
	models.DB.Save(&application)

	c.JSON(http.StatusOK, application)
}
