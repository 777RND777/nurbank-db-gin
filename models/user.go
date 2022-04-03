package models

type User struct {
	ID        int    `json:"id"`
	Password  string `json:"password" binding:"required"`
	PK        int    `json:"pk" gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Debt      int    `json:"debt"`
}

type CreateUserInput struct {
	ID        int    `json:"id"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Username  string `json:"username" binding:"required"`
}

type UpdateUserInput struct {
	ID       int    `json:"id"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}
