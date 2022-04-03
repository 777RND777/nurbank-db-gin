package controllers

import "time"

func GetCurrentTime() string {
	var t = time.Now()
	return t.Format("15:04:05 - 02/01/2006")
}
