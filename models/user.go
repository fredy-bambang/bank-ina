package models

import "time"

// User is a struct to hold user data
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"size:255" json:"name"`
	Email     string    `gorm:"size:100;unique" json:"email"`
	Password  string    `gorm:"size:100" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
