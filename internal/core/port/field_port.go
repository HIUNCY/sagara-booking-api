package port

import "github.com/HIUNCY/sagara-booking-api/internal/core/domain"

// DTO
type CreateFieldRequest struct {
	Name         string `json:"name"`
	PricePerHour int    `json:"price_per_hour"`
	Location     string `json:"location"`
}

type FieldRepository interface {
	Create(field *domain.Field) error
	GetAll() ([]domain.Field, error)
	GetByID(id uint) (*domain.Field, error)
	Update(field *domain.Field) error
	Delete(id uint) error
}

type FieldService interface {
	CreateField(req *CreateFieldRequest) error
	GetAllFields() ([]domain.Field, error)
	GetFieldByID(id uint) (*domain.Field, error)
	UpdateField(id uint, req *CreateFieldRequest) error
	DeleteField(id uint) error
}
