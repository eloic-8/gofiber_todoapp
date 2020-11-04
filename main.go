package main

import (
	"fmt"

	"github.com/eloicahhing/go-fiber-tutorial/database"
	"github.com/eloicahhing/go-fiber-tutorial/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"os"

	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())

	routes.Routes(app)

	// Load the Enviroment Setting
	envLoader()

	// Connection to Db
	database.ConnectDB()

	app.Listen(":" + os.Getenv("HOST_PORT"))
}

func envLoader() error {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return err
}
