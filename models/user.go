package models

type User struct {
	PK           int           `json:"pk" gorm:"primary_key"`
	ID           int           `json:"id"`
	FirstName    string        `json:"first_name"`
	LastName     string        `json:"last_name"`
	Username     string        `json:"username"`
	Nickname     string        `json:"nickname"`
	Debt         int           `json:"debt"`
	Applications []Application `gorm:"foreignKey:UserID"`
}

type CreateUserInput struct {
	ID        int    `json:"id" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
}

type UpdateUserInput struct {
	ID       int    `json:"id" binding:"required"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Debt     int    `json:"debt"`
}
