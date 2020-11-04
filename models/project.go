package models

import "time"

// Project Model
type Project struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"varchar(50),unique" json:"name"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	IsActive    bool      `json:"isActive"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
