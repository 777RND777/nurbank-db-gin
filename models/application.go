package models

type Application struct {
	ID          int    `json:"id" gorm:"unique"`
	UserID      int    `json:"user_id" binding:"required"`
	PK          int    `json:"pk" gorm:"primary_key"`
	Value       int    `json:"value"`
	RequestDate string `json:"request_date"`
	AnswerDate  string `json:"answer_date"`
	Approved    bool   `json:"approved"`
	IsAdmin     bool   `json:"is_admin"`
}

type CreateApplicationInput struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
	Value    int    `json:"value" binding:"required"`
	IsAdmin  bool   `json:"is_admin"`
}
