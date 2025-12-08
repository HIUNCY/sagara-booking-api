package repository

import (
	"github.com/HIUNCY/sagara-booking-api/internal/core/domain"
	"github.com/HIUNCY/sagara-booking-api/internal/core/port"
	"gorm.io/gorm"
)

type FieldRepositoryDB struct {
	db *gorm.DB
}

func NewFieldRepository(db *gorm.DB) port.FieldRepository {
	return &FieldRepositoryDB{db: db}
}

func (r *FieldRepositoryDB) Create(field *domain.Field) error {
	return r.db.Create(field).Error
}

func (r *FieldRepositoryDB) GetAll() ([]domain.Field, error) {
	var fields []domain.Field
	err := r.db.Find(&fields).Error
	return fields, err
}

func (r *FieldRepositoryDB) GetByID(id uint) (*domain.Field, error) {
	var field domain.Field
	err := r.db.First(&field, id).Error
	if err != nil {
		return nil, err
	}
	return &field, nil
}

func (r *FieldRepositoryDB) Update(field *domain.Field) error {
	return r.db.Save(field).Error
}

func (r *FieldRepositoryDB) Delete(id uint) error {
	return r.db.Delete(&domain.Field{}, id).Error
}
