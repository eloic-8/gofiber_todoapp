package models

import "time"

// Todo Model
type Todo struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Todos array
var Todos = []Todo{}