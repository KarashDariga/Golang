package models

type Article struct {
	ID      uint64 `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
