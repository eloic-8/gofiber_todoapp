package services

import (
	"github.com/eloicahhing/go-fiber-tutorial/database"
	"github.com/eloicahhing/go-fiber-tutorial/models"
)

// ReadAllTodos retrieve all todo
func ReadAllTodos() ([]models.Todo, error) {
	var todos []models.Todo
	result := database.DB.Find(&todos)

	return todos, result.Error
}

// ReadAllOpenTodos retrieve all open todo
func ReadAllOpenTodos() ([]models.Todo, error) {
	var todos []models.Todo
	result := database.DB.Find(&todos)

	return todos, result.Error
}
