package routes

import (
	"github.com/eloicahhing/go-fiber-tutorial/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Routes Defining all the routes of the API
func Routes(app *fiber.App) {

	app.Get("/", logger.New(), func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World 👋!")
	})

	// Route for todos
	todo := app.Group("/todos", logger.New())
	todo.Get("/", controllers.ReadAllTodo)
	todo.Post("/", controllers.CreateTodo)
	todo.Get("/:id", controllers.ReadTodo)
	todo.Put("/:id", controllers.UpdateTodo)
	todo.Delete("/:id", controllers.DeleteTodo)

}
