package handler

import (
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/HIUNCY/sagara-booking-api/internal/core/port"
    "github.com/gofiber/fiber/v2"
)

type mockUserService struct {
    registerErr error
    loginResp   *port.LoginResponse
    loginErr    error
}

func (m *mockUserService) Register(req *port.RegisterRequest) error { return m.registerErr }
func (m *mockUserService) Login(req *port.LoginRequest) (*port.LoginResponse, error) {
    if m.loginErr != nil { return nil, m.loginErr }
    return m.loginResp, nil
}

func TestUserHandler_Register(t *testing.T) {
    app := fiber.New()
    h := NewUserHandler(&mockUserService{})
    app.Post("/register", h.Register)

    body := map[string]any{"name": "A", "email": "a@mail", "password": "x"}
    b, _ := json.Marshal(body)
    req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(b))
    req.Header.Set("Content-Type", "application/json")
    resp, _ := app.Test(req)
    if resp.StatusCode != http.StatusCreated {
        t.Fatalf("expected 201, got %d", resp.StatusCode)
    }

    // invalid json
    req2 := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader([]byte("{")))
    req2.Header.Set("Content-Type", "application/json")
    resp2, _ := app.Test(req2)
    if resp2.StatusCode != http.StatusBadRequest {
        t.Fatalf("expected 400, got %d", resp2.StatusCode)
    }

    // service error
    app2 := fiber.New()
    h2 := NewUserHandler(&mockUserService{registerErr: errors.New("boom")})
    app2.Post("/register", h2.Register)
    req3 := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(b))
    req3.Header.Set("Content-Type", "application/json")
    resp3, _ := app2.Test(req3)
    if resp3.StatusCode != http.StatusInternalServerError {
        t.Fatalf("expected 500, got %d", resp3.StatusCode)
    }
}

func TestUserHandler_Login(t *testing.T) {
    app := fiber.New()
    h := NewUserHandler(&mockUserService{loginResp: &port.LoginResponse{Token: "tok"}})
    app.Post("/login", h.Login)

    body := map[string]any{"email": "a@mail", "password": "x"}
    b, _ := json.Marshal(body)
    req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
    req.Header.Set("Content-Type", "application/json")
    resp, _ := app.Test(req)
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("expected 200, got %d", resp.StatusCode)
    }

    // invalid json
    req2 := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader([]byte("{")))
    req2.Header.Set("Content-Type", "application/json")
    resp2, _ := app.Test(req2)
    if resp2.StatusCode != http.StatusBadRequest {
        t.Fatalf("expected 400, got %d", resp2.StatusCode)
    }

    // invalid credentials
    app2 := fiber.New()
    h2 := NewUserHandler(&mockUserService{loginErr: errors.New("bad")})
    app2.Post("/login", h2.Login)
    req3 := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
    req3.Header.Set("Content-Type", "application/json")
    resp3, _ := app2.Test(req3)
    if resp3.StatusCode != http.StatusUnauthorized {
        t.Fatalf("expected 401, got %d", resp3.StatusCode)
    }
}
