package handler

import (
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/HIUNCY/sagara-booking-api/internal/core/domain"
    "github.com/HIUNCY/sagara-booking-api/internal/core/port"
    "github.com/gofiber/fiber/v2"
)

type mockFieldService struct {
    createErr error
    fields    []domain.Field
    allErr    error
    byID      *domain.Field
    byIDErr   error
    updateErr error
    deleteErr error
}

func (m *mockFieldService) CreateField(req *port.CreateFieldRequest) error { return m.createErr }
func (m *mockFieldService) GetAllFields() ([]domain.Field, error) {
    if m.allErr != nil { return nil, m.allErr }
    return m.fields, nil
}
func (m *mockFieldService) GetFieldByID(id uint) (*domain.Field, error) {
    if m.byIDErr != nil { return nil, m.byIDErr }
    return m.byID, nil
}
func (m *mockFieldService) UpdateField(id uint, req *port.CreateFieldRequest) error { return m.updateErr }
func (m *mockFieldService) DeleteField(id uint) error { return m.deleteErr }

func TestFieldHandler_Create_And_GetAll(t *testing.T) {
    app := fiber.New()
    h := NewFieldHandler(&mockFieldService{fields: []domain.Field{{Name: "A"}}})
    app.Post("/fields", h.Create)
    app.Get("/fields", h.GetAll)

    // invalid json
    req := httptest.NewRequest(http.MethodPost, "/fields", bytes.NewReader([]byte("{")))
    req.Header.Set("Content-Type", "application/json")
    resp, _ := app.Test(req)
    if resp.StatusCode != http.StatusBadRequest {
        t.Fatalf("expected 400, got %d", resp.StatusCode)
    }

    // success create
    body := map[string]any{"name": "A", "price_per_hour": 10, "location": "L"}
    b, _ := json.Marshal(body)
    req2 := httptest.NewRequest(http.MethodPost, "/fields", bytes.NewReader(b))
    req2.Header.Set("Content-Type", "application/json")
    resp2, _ := app.Test(req2)
    if resp2.StatusCode != http.StatusCreated {
        t.Fatalf("expected 201, got %d", resp2.StatusCode)
    }

    // get all success
    req3 := httptest.NewRequest(http.MethodGet, "/fields", nil)
    resp3, _ := app.Test(req3)
    if resp3.StatusCode != http.StatusOK {
        t.Fatalf("expected 200, got %d", resp3.StatusCode)
    }

    // get all error
    app2 := fiber.New()
    h2 := NewFieldHandler(&mockFieldService{allErr: errors.New("boom")})
    app2.Get("/fields", h2.GetAll)
    req4 := httptest.NewRequest(http.MethodGet, "/fields", nil)
    resp4, _ := app2.Test(req4)
    if resp4.StatusCode != http.StatusInternalServerError {
        t.Fatalf("expected 500, got %d", resp4.StatusCode)
    }
}

func TestFieldHandler_GetByID_Update_Delete(t *testing.T) {
    app := fiber.New()
    h := NewFieldHandler(&mockFieldService{byID: &domain.Field{Name: "A"}})
    app.Get("/fields/:id", h.GetByID)
    app.Put("/fields/:id", h.Update)
    app.Delete("/fields/:id", h.Delete)

    // get by id success
    req := httptest.NewRequest(http.MethodGet, "/fields/1", nil)
    resp, _ := app.Test(req)
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("expected 200, got %d", resp.StatusCode)
    }

    // get by id not found
    app2 := fiber.New()
    h2 := NewFieldHandler(&mockFieldService{byIDErr: errors.New("not found")})
    app2.Get("/fields/:id", h2.GetByID)
    req2 := httptest.NewRequest(http.MethodGet, "/fields/1", nil)
    resp2, _ := app2.Test(req2)
    if resp2.StatusCode != http.StatusNotFound {
        t.Fatalf("expected 404, got %d", resp2.StatusCode)
    }

    // update invalid json
    req3 := httptest.NewRequest(http.MethodPut, "/fields/1", bytes.NewReader([]byte("{")))
    req3.Header.Set("Content-Type", "application/json")
    resp3, _ := app.Test(req3)
    if resp3.StatusCode != http.StatusBadRequest {
        t.Fatalf("expected 400, got %d", resp3.StatusCode)
    }

    // update success
    body := map[string]any{"name": "B", "price_per_hour": 20, "location": "X"}
    b, _ := json.Marshal(body)
    req4 := httptest.NewRequest(http.MethodPut, "/fields/1", bytes.NewReader(b))
    req4.Header.Set("Content-Type", "application/json")
    resp4, _ := app.Test(req4)
    if resp4.StatusCode != http.StatusOK {
        t.Fatalf("expected 200, got %d", resp4.StatusCode)
    }

    // delete success
    req5 := httptest.NewRequest(http.MethodDelete, "/fields/1", nil)
    resp5, _ := app.Test(req5)
    if resp5.StatusCode != http.StatusOK {
        t.Fatalf("expected 200, got %d", resp5.StatusCode)
    }
}
