package main

import (
	"github.com/sirupsen/logrus"
	"simpledrive/database"
	"simpledrive/routes"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		logrus.Println("No .env file found")
	}
	logrus.Println("Environment variables successfully loaded. Starting application...")
}

func main() {
	app := fiber.New()

	//Connect Database
	database.Connect()

	//Setup routes
	routes.Setup(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Simpledrive app set")
	})
	//Activate CORS
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	// Get the PORT from host env
	port := os.Getenv("PORT")

	if port == "" {
		// local development port
		port = "8000"
	}

	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
