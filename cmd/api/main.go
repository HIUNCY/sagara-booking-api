package main

import (
	"log"

	"github.com/HIUNCY/sagara-booking-api/internal/handler"
	"github.com/HIUNCY/sagara-booking-api/internal/repository"
	"github.com/HIUNCY/sagara-booking-api/internal/service"
	"github.com/HIUNCY/sagara-booking-api/pkg/database"
	"github.com/HIUNCY/sagara-booking-api/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	log.Fatal(app.Listen(":3000"))
}
