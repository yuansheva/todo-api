package models

type Todo struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Brand    string `json:"brand"`
	Platform string `json:"platform"`
	DueDate  string `json:"due_date"`
	Payment  int    `json:"payment"`
	Status   string `json:"status"`
}
