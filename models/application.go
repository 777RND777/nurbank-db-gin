package models

type UsersApplication struct {
	ID            int `json:"id"`
	UserID        int `json:"user_id" binding:"required"`
	ApplicationID int `json:"application_id" binding:"required"`
}

type Application struct {
	PK          int    `json:"pk"`
	ID          int    `json:"id" binding:"required"`
	UserID      int    `json:"user_id" binding:"required"`
	Value       int    `json:"value" binding:"required"`
	RequestDate string `json:"request_date" binding:"required"`
	AnswerDate  string `json:"answer_date"`
	Approved    bool   `json:"approved"`
	IsAdmin     bool   `json:"is_admin"`
}
