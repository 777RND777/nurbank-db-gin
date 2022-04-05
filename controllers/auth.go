package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"nur-bank-gin-api/models"
)

//TODO messages

type applicationAuth struct {
	UserID   int    `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CheckUser(userID int, password string) bool {
	var user models.User
	if err := models.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return false
	}

	return CheckPasswordHash(password, user.Password)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
