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

	if !CheckUser(input.UserID, input.Password) {
		c.JSON(http.StatusUnauthorized, models.Application{})
		return
	}

	var application = models.Application{
		ID:          GetUniqueID(),
		UserID:      input.UserID,
		Value:       input.Value,
		RequestDate: GetCurrentTime(),
		IsAdmin:     input.IsAdmin,
	}

	models.DB.Create(&application)

	c.JSON(http.StatusOK, application)
}

func GetApplication(c *gin.Context) {
	var application models.Application
	if err := models.DB.Where("id = ?", c.Param("id")).First(&application).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Application{})
		return
	}

	c.JSON(http.StatusOK, application)
}

func ApproveApplication(c *gin.Context) {
	var input applicationAuth
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Application{})
		return
	}

	if !CheckUser(input.UserID, input.Password) {
		c.JSON(http.StatusUnauthorized, models.Application{})
		return
	}

	var application models.Application
	if err := models.DB.Where("id = ?", c.Param("id")).First(&application).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Application{})
		return
	}
	if len(application.AnswerDate) > 0 {
		c.JSON(http.StatusNoContent, models.Application{})
		return
	}

	application.AnswerDate = GetCurrentTime()
	application.Approved = true
	models.DB.Save(&application)

	var user models.User
	//no check because of CheckUser func
	models.DB.Where("id = ?", application.UserID).First(&user)

	user.Debt += application.Value
	models.DB.Save(&user)

	c.JSON(http.StatusOK, application)
}

func DeclineApplication(c *gin.Context) {
	var input applicationAuth
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Application{})
		return
	}

	if !CheckUser(input.UserID, input.Password) {
		c.JSON(http.StatusUnauthorized, models.Application{})
		return
	}

	var application models.Application
	if err := models.DB.Where("id = ?", c.Param("id")).First(&application).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Application{})
		return
	}
	if len(application.AnswerDate) > 0 {
		c.JSON(http.StatusNoContent, models.Application{})
		return
	}

	application.AnswerDate = GetCurrentTime()
	models.DB.Save(&application)

	c.JSON(http.StatusOK, application)
}
