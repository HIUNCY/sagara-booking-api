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

// CreateBooking godoc
// @Summary      Create a new booking
// @Description  Book a field. Checks for schedule overlap.
// @Tags         Bookings
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        booking body port.BookingRequest true "Booking Data"
// @Success      201 {object} port.DataResponse
// @Failure      400 {object} port.ErrorResponse "Invalid Input / Overlap"
// @Failure      401 {object} port.ErrorResponse "Unauthorized"
// @Router       /bookings [post]
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

// GetAllBookings godoc
// @Summary      Get all bookings history
// @Description  Retrieve a list of all bookings (Admin/User).
// @Tags         Bookings
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} port.DataResponse
// @Failure      500 {object} port.ErrorResponse
// @Router       /bookings [get]
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

// GetBookingByID godoc
// @Summary      Get booking details
// @Description  Get detailed information about a specific booking by ID.
// @Tags         Bookings
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Booking ID"
// @Success      200 {object} port.DataResponse
// @Failure      404 {object} port.ErrorResponse
// @Router       /bookings/{id} [get]
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

// PayBooking godoc
// @Summary      Pay for a booking (Mock Payment)
// @Description  Change booking status from pending to paid.
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        payment body object{booking_id=int} true "JSON: {booking_id: 1}"
// @Success      200 {object} port.MessageResponse
// @Failure      400 {object} port.ErrorResponse
// @Router       /payments [post]
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
