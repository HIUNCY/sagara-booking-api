package handler

import (
	"github.com/HIUNCY/sagara-booking-api/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service port.UserService
}

func NewUserHandler(service port.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Register godoc
// @Summary      Register New User
// @Description  Create a new user account. Role can be 'user' or 'admin'.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user body port.RegisterRequest true "User Data"
// @Success      201 {object} port.MessageResponse "message: User created successfully"
// @Failure      400 {object} port.ErrorResponse "Invalid Input"
// @Failure      500 {object} port.ErrorResponse "Internal Server Error"
// @Router       /register [post]
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req port.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Input"})
	}

	if err := h.service.Register(&req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User created successfully"})
}

// Login godoc
// @Summary      User Login
// @Description  Authenticate user and return JWT token.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        credentials body port.LoginRequest true "Login Credentials"
// @Success      200 {object} port.LoginResponse
// @Failure      400 {object} port.ErrorResponse "Invalid Input"
// @Failure      401 {object} port.ErrorResponse "Invalid Email or Password"
// @Router       /login [post]
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req port.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Input"})
	}

	res, err := h.service.Login(&req)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(res)
}
