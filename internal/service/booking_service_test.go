package service

import (
    "errors"
    "testing"
    "time"

    "github.com/HIUNCY/sagara-booking-api/internal/core/domain"
    "github.com/HIUNCY/sagara-booking-api/internal/core/port"
)

type mockBookingRepo struct {
    created []*domain.Booking
    avail map[uint]bool
    availErr error
    byID map[uint]*domain.Booking
    updateErr error
}

func (m *mockBookingRepo) Create(b *domain.Booking) error {
    if m.created == nil {
        m.created = []*domain.Booking{}
    }
    b.ID = uint(len(m.created) + 1)
    m.created = append(m.created, b)
    if m.byID == nil {
        m.byID = map[uint]*domain.Booking{}
    }
    m.byID[b.ID] = b
    return nil
}

func (m *mockBookingRepo) CheckAvailability(fieldID uint, start, end time.Time) (bool, error) {
    if m.availErr != nil {
        return false, m.availErr
    }
    if m.avail != nil {
        if v, ok := m.avail[fieldID]; ok {
            return v, nil
        }
    }
    return false, nil
}

func (m *mockBookingRepo) GetByID(id uint) (*domain.Booking, error) {
    if b, ok := m.byID[id]; ok {
        return b, nil
    }
    return nil, errors.New("not found")
}

func (m *mockBookingRepo) UpdateStatus(id uint, status string) error {
    if m.updateErr != nil {
        return m.updateErr
    }
    if b, ok := m.byID[id]; ok {
        b.Status = status
        return nil
    }
    return errors.New("not found")
}

func (m *mockBookingRepo) GetAll() ([]domain.Booking, error) {
    res := make([]domain.Booking, 0, len(m.byID))
    for _, b := range m.byID {
        res = append(res, *b)
    }
    return res, nil
}

func TestBookingService_CreateBooking_ValidationsAndSuccess(t *testing.T) {
    repo := &mockBookingRepo{avail: map[uint]bool{1: false}}
    svc := NewBookingService(repo)
    start := time.Now().Add(time.Hour)
    end := start.Add(time.Hour)

    // invalid time
    if _, err := svc.CreateBooking(10, &port.BookingRequest{FieldID: 1, StartTime: end, EndTime: start}); err == nil {
        t.Fatalf("expected error for end before start")
    }

    // overlap
    repo.avail[1] = true
    if _, err := svc.CreateBooking(10, &port.BookingRequest{FieldID: 1, StartTime: start, EndTime: end}); err == nil {
        t.Fatalf("expected overlap error")
    }

    // success
    repo.avail[1] = false
    b, err := svc.CreateBooking(10, &port.BookingRequest{FieldID: 1, StartTime: start, EndTime: end});
    if err != nil || b == nil {
        t.Fatalf("expected booking created, got err=%v", err)
    }
    if b.Status != "pending" || b.UserID != 10 || b.FieldID != 1 {
        t.Fatalf("unexpected booking values: %+v", b)
    }
}

func TestBookingService_GetAndPay(t *testing.T) {
    repo := &mockBookingRepo{}
    svc := NewBookingService(repo)
    start := time.Now().Add(time.Hour)
    end := start.Add(time.Hour)

    b, _ := svc.CreateBooking(2, &port.BookingRequest{FieldID: 3, StartTime: start, EndTime: end})
    list, _ := svc.GetAllBookings()
    if len(list) != 1 {
        t.Fatalf("expected 1 booking, got %d", len(list))
    }
    got, _ := svc.GetBookingByID(b.ID)
    if got.ID != b.ID {
        t.Fatalf("expected same booking id")
    }
    if err := svc.PayBooking(b.ID); err != nil {
        t.Fatalf("pay error: %v", err)
    }
    if b.Status != "paid" {
        t.Fatalf("expected status paid, got %s", b.Status)
    }
}
