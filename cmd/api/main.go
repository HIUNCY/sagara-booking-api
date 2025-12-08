package main

import (
	"log"

	"github.com/HIUNCY/sagara-booking-api/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Database error: %v", err)
	}
	_ = db

	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Sagara Clean Architecture API is Ready!")
	})

	log.Fatal(app.Listen(":3000"))
}
