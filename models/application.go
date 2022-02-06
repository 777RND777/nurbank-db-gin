package models

type Application struct {
	PK          int    `json:"pk"`
	ID          int    `json:"id" binding:"required"`
	UserID      int    `json:"user_id" binding:"required"`
	User        User   `gorm:"foreignKey:UserID"`
	Value       int    `json:"value" binding:"required"`
	RequestDate string `json:"request_date" binding:"required"`
	AnswerDate  string `json:"answer_date"`
	Approved    bool   `json:"approved"`
	IsAdmin     bool   `json:"is_admin"`
}
