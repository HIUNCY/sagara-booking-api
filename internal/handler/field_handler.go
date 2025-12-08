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

// CreateField godoc
// @Summary      Create New Field (Admin Only)
// @Description  Add a new sports field to the system. Requires Admin role.
// @Tags         Fields
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        field body port.CreateFieldRequest true "Field Data"
// @Success      201 {object} port.MessageResponse
// @Failure      400 {object} port.ErrorResponse "Invalid Input"
// @Failure      403 {object} port.ErrorResponse "Forbidden: Admin only"
// @Failure      500 {object} port.ErrorResponse
// @Router       /fields [post]
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

// GetAllFields godoc
// @Summary      Get All Fields
// @Description  Retrieve a list of all available sports fields.
// @Tags         Fields
// @Produce      json
// @Security     BearerAuth
// @Success      200 {array} port.FieldResponse
// @Failure      500 {object} port.ErrorResponse
// @Router       /fields [get]
func (h *FieldHandler) GetAll(c *fiber.Ctx) error {
	fields, err := h.service.GetAllFields()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fields)
}

// GetFieldByID godoc
// @Summary      Get Field Detail
// @Description  Get detailed information of a specific field by ID.
// @Tags         Fields
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Field ID"
// @Success      200 {object} port.FieldResponse
// @Failure      404 {object} port.ErrorResponse
// @Router       /fields/{id} [get]
func (h *FieldHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	field, err := h.service.GetFieldByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Field not found"})
	}
	return c.JSON(field)
}

// UpdateField godoc
// @Summary      Update Field (Admin Only)
// @Description  Update existing field data. Requires Admin role.
// @Tags         Fields
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Field ID"
// @Param        field body port.CreateFieldRequest true "Updated Data"
// @Success      200 {object} port.MessageResponse
// @Failure      400 {object} port.ErrorResponse
// @Failure      403 {object} port.ErrorResponse "Forbidden"
// @Failure      500 {object} port.ErrorResponse
// @Router       /fields/{id} [put]
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

// DeleteField godoc
// @Summary      Delete Field (Admin Only)
// @Description  Permanently remove a field. Requires Admin role.
// @Tags         Fields
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Field ID"
// @Success      200 {object} port.MessageResponse
// @Failure      403 {object} port.ErrorResponse "Forbidden"
// @Failure      500 {object} port.ErrorResponse
// @Router       /fields/{id} [delete]
func (h *FieldHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.service.DeleteField(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Field deleted successfully"})
}
