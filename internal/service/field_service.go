package service

import (
	"github.com/HIUNCY/sagara-booking-api/internal/core/domain"
	"github.com/HIUNCY/sagara-booking-api/internal/core/port"
)

type FieldServiceImpl struct {
	repo port.FieldRepository
}

func NewFieldService(repo port.FieldRepository) port.FieldService {
	return &FieldServiceImpl{repo: repo}
}

func (s *FieldServiceImpl) CreateField(req *port.CreateFieldRequest) error {
	field := &domain.Field{
		Name:         req.Name,
		PricePerHour: req.PricePerHour,
		Location:     req.Location,
	}
	return s.repo.Create(field)
}

func (s *FieldServiceImpl) GetAllFields() ([]domain.Field, error) {
	return s.repo.GetAll()
}

func (s *FieldServiceImpl) GetFieldByID(id uint) (*domain.Field, error) {
	return s.repo.GetByID(id)
}

func (s *FieldServiceImpl) UpdateField(id uint, req *port.CreateFieldRequest) error {
	field, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	field.Name = req.Name
	field.PricePerHour = req.PricePerHour
	field.Location = req.Location

	return s.repo.Update(field)
}

func (s *FieldServiceImpl) DeleteField(id uint) error {
	return s.repo.Delete(id)
}
