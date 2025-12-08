package main

import (
	"log"
	"os"

	_ "github.com/HIUNCY/sagara-booking-api/docs"
	"github.com/HIUNCY/sagara-booking-api/internal/handler"
	"github.com/HIUNCY/sagara-booking-api/internal/repository"
	"github.com/HIUNCY/sagara-booking-api/internal/service"
	"github.com/HIUNCY/sagara-booking-api/pkg/database"
	"github.com/HIUNCY/sagara-booking-api/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	swagger "github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

// @title           Sagara Booking API
// @version         1.0
// @description     Backend Developer Take Home Test - Booking System.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.email  support@sagara.id
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host            sagara-booking-api-f264e78236b6.herokuapp.com
// @BasePath        /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, relying on Heroku Config Vars.")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Database error: %v", err)
	}

	// USER FEATURE
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// FIELD FEATURE
	fieldRepo := repository.NewFieldRepository(db)
	fieldService := service.NewFieldService(fieldRepo)
	fieldHandler := handler.NewFieldHandler(fieldService)

	// BOOKING FEATURE
	bookingRepo := repository.NewBookingRepository(db)
	bookingService := service.NewBookingService(bookingRepo)
	bookingHandler := handler.NewBookingHandler(bookingService)

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Sagara Backend Test API is Running!")
	})

	// SWAGGER ROUTES
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "/swagger/doc.json",
	}))

	api := app.Group("/api")

	// AUTH ROUTES
	api.Post("/register", userHandler.Register)
	api.Post("/login", userHandler.Login)

	// FIELD ROUTES
	fields := api.Group("/fields", middleware.Protected)
	fields.Get("/", fieldHandler.GetAll)
	fields.Get("/:id", fieldHandler.GetByID)
	fields.Post("/", middleware.AdminOnly, fieldHandler.Create)
	fields.Put("/:id", middleware.AdminOnly, fieldHandler.Update)
	fields.Delete("/:id", middleware.AdminOnly, fieldHandler.Delete)

	// BOOKING AND PAYMENT ROUTES
	bookings := api.Group("/bookings", middleware.Protected)
	bookings.Get("/", bookingHandler.GetAll)
	bookings.Get("/:id", bookingHandler.GetByID)
	bookings.Post("/", bookingHandler.Create)
	api.Post("/payments", middleware.Protected, bookingHandler.Pay)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
