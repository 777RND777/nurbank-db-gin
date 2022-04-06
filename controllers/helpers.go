package controllers

import (
	"math/rand"
	"time"

	"nur-bank-gin-api/models"
)

func GetCurrentTime() string {
	var t = time.Now()
	return t.Format("15:04:05 - 02/01/2006")
}

func GetUniqueID() int {
	var application models.Application
	var id = rand.Intn(999999999-100000000) + 100000000
	var err = models.DB.Where("id = ?", id).First(&application).Error
	for err != nil {
		id = rand.Intn(999999999-100000000) + 100000000
		err = models.DB.Where("id = ?", id).First(&application).Error
	}
	return id
}
