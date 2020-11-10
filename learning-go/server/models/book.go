package models

// Book model for db
type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  uint   `json:"pages"`
}
