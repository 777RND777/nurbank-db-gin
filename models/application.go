package models

type UsersApplication struct {
	ID            int `json:"id" db:"id"`
	UserID        int `json:"user_id" db:"user_id" binding:"required"`
	ApplicationID int `json:"application_id" db:"application_id" binding:"required"`
}

type Application struct {
	PK          int    `json:"pk" db:"pk"`
	ID          int    `json:"id" db:"id" binding:"required"`
	UserID      int    `json:"user_id" db:"user_id" binding:"required"`
	Value       int    `json:"value" db:"value" binding:"required"`
	RequestDate string `json:"request_date" db:"request_date" binding:"required"`
	AnswerDate  string `json:"answer_date" db:"answer_date"`
	Approved    bool   `json:"approved" db:"approved"`
	IsAdmin     bool   `json:"is_admin" db:"is_admin"`
}
