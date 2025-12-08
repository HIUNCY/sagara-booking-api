package service

import (
	"errors"

	"github.com/HIUNCY/sagara-booking-api/internal/core/domain"
	"github.com/HIUNCY/sagara-booking-api/internal/core/port"
)

type BookingServiceImpl struct {
	repo port.BookingRepository
}

func NewBookingService(repo port.BookingRepository) port.BookingService {
	return &BookingServiceImpl{repo: repo}
}

func (s *BookingServiceImpl) CreateBooking(userID uint, req *port.BookingRequest) (*domain.Booking, error) {
	if req.EndTime.Before(req.StartTime) {
		return nil, errors.New("start time must be before end time")
	}

	isBooked, err := s.repo.CheckAvailability(req.FieldID, req.StartTime, req.EndTime)
	if err != nil {
		return nil, err
	}
	if isBooked {
		return nil, errors.New("field is already booked at this time")
	}

	booking := &domain.Booking{
		UserID:    userID,
		FieldID:   req.FieldID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Status:    "pending",
	}

	err = s.repo.Create(booking)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (s *BookingServiceImpl) GetAllBookings() ([]domain.Booking, error) {
	return s.repo.GetAll()
}

func (s *BookingServiceImpl) GetBookingByID(id uint) (*domain.Booking, error) {
	return s.repo.GetByID(id)
}

func (s *BookingServiceImpl) PayBooking(bookingID uint) error {
	return s.repo.UpdateStatus(bookingID, "paid")
}
