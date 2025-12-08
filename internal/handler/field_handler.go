package handler

import (
	"strconv"

	"github.com/HIUNCY/sagara-booking-api/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type FieldHandler struct {
	service port.FieldService
}

func NewFieldHandler(service port.FieldService) *FieldHandler {
	return &FieldHandler{service: service}
}

func (h *FieldHandler) Create(c *fiber.Ctx) error {
	var req port.CreateFieldRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Input"})
	}

	if err := h.service.CreateField(&req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "Field created successfully"})
}

func (h *FieldHandler) GetAll(c *fiber.Ctx) error {
	fields, err := h.service.GetAllFields()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fields)
}

func (h *FieldHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	field, err := h.service.GetFieldByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Field not found"})
	}
	return c.JSON(field)
}

func (h *FieldHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var req port.CreateFieldRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Input"})
	}

	if err := h.service.UpdateField(uint(id), &req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Field updated successfully"})
}

func (h *FieldHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.service.DeleteField(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Field deleted successfully"})
}
