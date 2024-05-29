package models

import "time"

type Author struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100)" json:"name"`
	Gender string `gorm:"type:char(1)" json:"gender"`
	Email string `gorm:"type:varchar(100)" json:"email"`
	Password string `gorm:"type:varchar(100)" json:"password"`
	Age int `gorm:"type:integer" json:"age"`
	CreatedAt time.Time `json:created_at`
	UpdatedAt time.Time `json:updated_at`
}

type Register struct{
	Name string `json:"name"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	Password string `json:"password"`
	Age int `json:"age"`
}

type Login struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type AuthorBookResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	Age int `json:"age"`
}