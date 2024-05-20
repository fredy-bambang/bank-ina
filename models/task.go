package models

import "time"

// Task is a struct to hold task data
type Task struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	UserID      uint      `gorm:"index not null" json:"user_id" binding:"required,number"`
	Title       string    `gorm:"size:255" json:"title" binding:"required,min=3,max=50"`
	Description string    `json:"description"`
	Status      string    `gorm:"size:50" json:"status" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
