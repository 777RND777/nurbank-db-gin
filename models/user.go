package models

type User struct {
	PK        int    `json:"-" gorm:"primary_key"`
	ID        int    `json:"id" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Debt      int    `json:"debt"`
}
