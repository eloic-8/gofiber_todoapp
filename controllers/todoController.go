package controllers

import (
	"strconv"

	"github.com/eloicahhing/go-fiber-tutorial/services"

	"github.com/eloicahhing/go-fiber-tutorial/database"
	"github.com/eloicahhing/go-fiber-tutorial/models"
	"github.com/gofiber/fiber/v2"
)

// CreateTodo create a new Todo
func CreateTodo(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json") // "application/json"

	Todo := models.Todo{
		Name:      ctx.FormValue("name"),
		Completed: false,
	}
	if response := database.DB.Create(&Todo); response.Error != nil {
		panic("An error occurred when storing the new role: " + response.Error.Error())
	}
	response := ctx.Status(fiber.StatusCreated).JSON(Todo)
	if response != nil {
		panic("Error occurred when returning JSON of a role: " + response.Error())
	}
	return response
}

// ReadTodo return Todo by id
func ReadTodo(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json") // "application/json"
	paramID := ctx.Params("id")
	if response := database.DB.Find(&models.Todos, paramID); response.Error != nil {
		panic("Error occurred while retrieving roles from the database: " + response.Error.Error())
	}
	response := ctx.Status(fiber.StatusOK).JSON(models.Todos)
	if response != nil {
		panic("Error occurred when returning JSON of roles: " + response.Error())
	}
	return response
}

// UpdateTodo update Todo by id
func UpdateTodo(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json") // "application/json"
	paramID := ctx.Params("id")
	int, err := strconv.Atoi(paramID)
	convCompleted, err := strconv.Atoi(ctx.FormValue("completed"))
	completed := false
	if err != nil {
		panic(err)
	}
	if convCompleted == 1 {
		completed = true
	}
	Todo := models.Todo{
		ID:        int,
		Name:      ctx.FormValue("name"),
		Completed: completed,
	}

	// if response := database.DB.Save(&Todo); response.Error != nil {
	if response := database.DB.Updates(&Todo); response.Error != nil {
		panic("Error occurred while retrieving roles from the database: " + response.Error.Error())
	}
	response := ctx.Status(fiber.StatusOK).JSON(models.Todos)
	if response != nil {
		panic("Error occurred when returning JSON of roles: " + response.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "the todo has been deleted",
	})
}

// DeleteTodo delete Todo by id
func DeleteTodo(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json") // "application/json"
	paramID := ctx.Params("id")
	if response := database.DB.Delete(&models.Todos, paramID); response.Error != nil {
		panic("Error occurred while retrieving roles from the database: " + response.Error.Error())
	}
	response := ctx.Status(fiber.StatusOK).JSON(models.Todos)
	if response != nil {
		panic("Error occurred when returning JSON of roles: " + response.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "the todo has been updated",
	})
}

// ReadAllTodo return all Todos
func ReadAllTodo(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json") // "application/json"

	response, err := services.ReadAllTodos()

	if err != nil {
		_ = ctx.JSON(&fiber.Map{
			"status": false,
			"error":  err,
		})
	}
	return ctx.JSON(&fiber.Map{
		"status": true,
		"books":  response,
	})
}
