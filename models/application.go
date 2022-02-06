package models

type Application struct {
	PK          int    `json:"pk"`
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	User        User   `json:"user" gorm:"foreignKey:UserID"`
	Value       int    `json:"value"`
	RequestDate string `json:"request_date"`
	AnswerDate  string `json:"answer_date"`
	Approved    bool   `json:"approved"`
	IsAdmin     bool   `json:"is_admin"`
}

type CreateApplicationInput struct {
	ID          int    `json:"id" binding:"required"`
	UserID      int    `json:"user_id" binding:"required"`
	Value       int    `json:"value" binding:"required"`
	RequestDate string `json:"request_date" binding:"required"`
	IsAdmin     bool   `json:"is_admin"`
}

type UpdateApplicationInput struct {
	ID         int    `json:"id" binding:"required"`
	AnswerDate string `json:"answer_date"`
	Approved   bool   `json:"approved"`
}
