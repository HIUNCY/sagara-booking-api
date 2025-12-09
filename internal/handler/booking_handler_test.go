package handler

import (
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"

    "github.com/HIUNCY/sagara-booking-api/internal/core/domain"
    "github.com/HIUNCY/sagara-booking-api/internal/core/port"
    "github.com/gofiber/fiber/v2"
)

type mockBookingService struct {
    createResp *domain.Booking
    createErr  error
    allResp    []domain.Booking
    allErr     error
    byIDResp   *domain.Booking
    byIDErr    error
    payErr     error
}

func (m *mockBookingService) CreateBooking(userID uint, req *port.BookingRequest) (*domain.Booking, error) {
    if m.createErr != nil { return nil, m.createErr }
    return m.createResp, nil
}
func (m *mockBookingService) PayBooking(bookingID uint) error { return m.payErr }
func (m *mockBookingService) GetAllBookings() ([]domain.Booking, error) {
    if m.allErr != nil { return nil, m.allErr }
    return m.allResp, nil
}
func (m *mockBookingService) GetBookingByID(id uint) (*domain.Booking, error) {
    if m.byIDErr != nil { return nil, m.byIDErr }
    return m.byIDResp, nil
}

func TestBookingHandler_Create_UnauthorizedAndSuccess(t *testing.T) {
    app := fiber.New()
    h := NewBookingHandler(&mockBookingService{createResp: &domain.Booking{}})
    app.Post("/bookings", func(c *fiber.Ctx) error { return h.Create(c) })

    // unauthorized (no locals)
    body := map[string]any{"field_id": 1, "start_time": time.Now(), "end_time": time.Now().Add(time.Hour)}
    b, _ := json.Marshal(body)
    req := httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewReader(b))
    req.Header.Set("Content-Type", "application/json")
    resp, _ := app.Test(req)
    if resp.StatusCode != http.StatusUnauthorized {
        t.Fatalf("expected 401, got %d", resp.StatusCode)
    }

    // success with user_id set
    app2 := fiber.New()
    app2.Post("/bookings", func(c *fiber.Ctx) error {
        c.Locals("user_id", float64(5))
        return h.Create(c)
    })
    req2 := httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewReader(b))
    req2.Header.Set("Content-Type", "application/json")
    resp2, _ := app2.Test(req2)
    if resp2.StatusCode != http.StatusCreated {
        t.Fatalf("expected 201, got %d", resp2.StatusCode)
    }

    // invalid json
    app3 := fiber.New()
    app3.Post("/bookings", func(c *fiber.Ctx) error { c.Locals("user_id", float64(1)); return h.Create(c) })
    req3 := httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewReader([]byte("{")))
    req3.Header.Set("Content-Type", "application/json")
    resp3, _ := app3.Test(req3)
    if resp3.StatusCode != http.StatusBadRequest {
        t.Fatalf("expected 400, got %d", resp3.StatusCode)
    }

    // service error
    app4 := fiber.New()
    hErr := NewBookingHandler(&mockBookingService{createErr: errors.New("overlap")})
    app4.Post("/bookings", func(c *fiber.Ctx) error { c.Locals("user_id", float64(1)); return hErr.Create(c) })
    req4 := httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewReader(b))
    req4.Header.Set("Content-Type", "application/json")
    resp4, _ := app4.Test(req4)
    if resp4.StatusCode != http.StatusBadRequest {
        t.Fatalf("expected 400 from service error, got %d", resp4.StatusCode)
    }
}

func TestBookingHandler_GetAll_And_GetByID(t *testing.T) {
    app := fiber.New()
    h := NewBookingHandler(&mockBookingService{allResp: []domain.Booking{{}}, byIDResp: &domain.Booking{}})
    app.Get("/bookings", h.GetAll)
    app.Get("/bookings/:id", h.GetByID)

    // get all success
    req := httptest.NewRequest(http.MethodGet, "/bookings", nil)
    resp, _ := app.Test(req)
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("expected 200, got %d", resp.StatusCode)
    }

    // get by id not found
    app2 := fiber.New()
    h2 := NewBookingHandler(&mockBookingService{byIDErr: errors.New("not found")})
    app2.Get("/bookings/:id", h2.GetByID)
    req2 := httptest.NewRequest(http.MethodGet, "/bookings/1", nil)
    resp2, _ := app2.Test(req2)
    if resp2.StatusCode != http.StatusNotFound {
        t.Fatalf("expected 404, got %d", resp2.StatusCode)
    }

    // get by id success
    req3 := httptest.NewRequest(http.MethodGet, "/bookings/2", nil)
    resp3, _ := app.Test(req3)
    if resp3.StatusCode != http.StatusOK {
        t.Fatalf("expected 200, got %d", resp3.StatusCode)
    }
}

func TestBookingHandler_Pay(t *testing.T) {
    app := fiber.New()
    h := NewBookingHandler(&mockBookingService{})
    app.Post("/payments", h.Pay)

    // invalid json
    req := httptest.NewRequest(http.MethodPost, "/payments", bytes.NewReader([]byte("{")))
    req.Header.Set("Content-Type", "application/json")
    resp, _ := app.Test(req)
    if resp.StatusCode != http.StatusBadRequest {
        t.Fatalf("expected 400, got %d", resp.StatusCode)
    }

    // success
    body := map[string]any{"booking_id": 1}
    b, _ := json.Marshal(body)
    req2 := httptest.NewRequest(http.MethodPost, "/payments", bytes.NewReader(b))
    req2.Header.Set("Content-Type", "application/json")
    resp2, _ := app.Test(req2)
    if resp2.StatusCode != http.StatusOK {
        t.Fatalf("expected 200, got %d", resp2.StatusCode)
    }

    // service error
    app2 := fiber.New()
    hErr := NewBookingHandler(&mockBookingService{payErr: errors.New("boom")})
    app2.Post("/payments", hErr.Pay)
    req3 := httptest.NewRequest(http.MethodPost, "/payments", bytes.NewReader(b))
    req3.Header.Set("Content-Type", "application/json")
    resp3, _ := app2.Test(req3)
    if resp3.StatusCode != http.StatusInternalServerError {
        t.Fatalf("expected 500, got %d", resp3.StatusCode)
    }
}
