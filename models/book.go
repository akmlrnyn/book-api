package models

import "time"

type Book struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	AuthorID uint `json:"author_id"`
	Description string `json:"description"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}