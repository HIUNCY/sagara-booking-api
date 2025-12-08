package repository

import (
	"time"

	"github.com/HIUNCY/sagara-booking-api/internal/core/domain"
	"github.com/HIUNCY/sagara-booking-api/internal/core/port"
	"gorm.io/gorm"
)

type BookingRepositoryDB struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) port.BookingRepository {
	return &BookingRepositoryDB{db: db}
}

func (r *BookingRepositoryDB) Create(booking *domain.Booking) error {
	return r.db.Create(booking).Error
}

func (r *BookingRepositoryDB) CheckAvailability(fieldID uint, start, end time.Time) (bool, error) {
	var count int64
	err := r.db.Model(&domain.Booking{}).
		Where("field_id = ?", fieldID).
		Where("status != ?", "cancelled").
		Where("start_time < ? AND end_time > ?", end, start).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *BookingRepositoryDB) GetAll() ([]domain.Booking, error) {
	var bookings []domain.Booking
	err := r.db.Preload("User").Preload("Field").Order("created_at desc").Find(&bookings).Error
	return bookings, err
}

func (r *BookingRepositoryDB) GetByID(id uint) (*domain.Booking, error) {
	var booking domain.Booking
	err := r.db.Preload("User").Preload("Field").First(&booking, id).Error
	return &booking, err
}

func (r *BookingRepositoryDB) UpdateStatus(id uint, status string) error {
	return r.db.Model(&domain.Booking{}).Where("id = ?", id).Update("status", status).Error
}
