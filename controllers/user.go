package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nur-bank-gin-api/models"
)

func CreateUser(c *gin.Context) {
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.User{})
		return
	}

	var user models.User
	if err := models.DB.Where("id = ?", input.ID).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, models.User{})
		return
	}

	var passwordHash, err = HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.User{})
		return
	}

	user = models.User{
		ID:        input.ID,
		Password:  passwordHash,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Username:  input.Username,
		Nickname:  input.Username,
	}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, user)
}

func GetUserList(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, models.User{})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserApplications(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, []models.Application{})
		return
	}

	var applications []models.Application
	models.DB.Find(&applications).Where("user_id = ?", c.Param("id"))

	c.JSON(http.StatusOK, applications)
}

func GetUserPending(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("user_id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Application{})
		return
	}

	var applications []models.Application
	models.DB.Find(&applications).Where("user_id = ?", c.Param("id"))
	if len(applications) == 0 || len(applications[len(applications)-1].AnswerDate) > 0 {
		c.JSON(http.StatusNoContent, models.Application{})
		return
	}

	c.JSON(http.StatusOK, applications[len(applications)-1])
}

func UpdateUser(c *gin.Context) {
	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.User{})
		return
	}

	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, models.User{})
		return
	}

	if !CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, models.User{})
		return
	}

	input.Password = user.Password //avoid resetting password
	models.DB.Model(&user).Update(input)

	c.JSON(http.StatusCreated, user)
}
