package port

import (
	"time"

	"github.com/HIUNCY/sagara-booking-api/internal/core/domain"
)

type BookingRequest struct {
	FieldID   uint      `json:"field_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type BookingRepository interface {
	Create(booking *domain.Booking) error
	CheckAvailability(fieldID uint, start, end time.Time) (bool, error)
	GetByID(id uint) (*domain.Booking, error)
	UpdateStatus(id uint, status string) error
	GetAll() ([]domain.Booking, error)
}

type BookingService interface {
	CreateBooking(userID uint, req *BookingRequest) (*domain.Booking, error)
	PayBooking(bookingID uint) error
	GetAllBookings() ([]domain.Booking, error)
	GetBookingByID(id uint) (*domain.Booking, error)
}
