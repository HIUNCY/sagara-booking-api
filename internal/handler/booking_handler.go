package handler

import (
	"strconv"

	"github.com/HIUNCY/sagara-booking-api/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type BookingHandler struct {
	service port.BookingService
}

func NewBookingHandler(service port.BookingService) *BookingHandler {
	return &BookingHandler{service: service}
}

func (h *BookingHandler) Create(c *fiber.Ctx) error {
	userIDFloat, ok := c.Locals("user_id").(float64)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}
	userID := uint(userIDFloat)

	var req port.BookingRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Input format"})
	}

	booking, err := h.service.CreateBooking(userID, &req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Booking created successfully",
		"data":    booking,
	})
}

func (h *BookingHandler) GetAll(c *fiber.Ctx) error {
	bookings, err := h.service.GetAllBookings()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{
		"message": "Success retrieving bookings",
		"data":    bookings,
	})
}

func (h *BookingHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	booking, err := h.service.GetBookingByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Booking not found"})
	}

	return c.JSON(fiber.Map{
		"message": "Success retrieving booking detail",
		"data":    booking,
	})
}

func (h *BookingHandler) Pay(c *fiber.Ctx) error {
	var req struct {
		BookingID uint `json:"booking_id"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Input"})
	}

	if err := h.service.PayBooking(req.BookingID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Payment successful, booking status updated to paid"})
}
