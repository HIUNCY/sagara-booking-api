package service

import (
    "errors"
    "os"
    "testing"

    "github.com/HIUNCY/sagara-booking-api/internal/core/domain"
    "github.com/HIUNCY/sagara-booking-api/internal/core/port"
)

type mockUserRepo struct {
    users map[string]*domain.User
    createErr error
    getErrByEmail map[string]error
}

func (m *mockUserRepo) CreateUser(user *domain.User) error {
    if m.createErr != nil {
        return m.createErr
    }
    if m.users == nil {
        m.users = map[string]*domain.User{}
    }
    // simulate auto ID
    user.ID = uint(len(m.users) + 1)
    m.users[user.Email] = user
    return nil
}

func (m *mockUserRepo) GetByEmail(email string) (*domain.User, error) {
    if m.getErrByEmail != nil {
        if err, ok := m.getErrByEmail[email]; ok {
            return nil, err
        }
    }
    u, ok := m.users[email]
    if !ok {
        return nil, errors.New("not found")
    }
    return u, nil
}

func TestUserService_Register_DefaultRoleAndHash(t *testing.T) {
    repo := &mockUserRepo{}
    svc := NewUserService(repo)
    req := &port.RegisterRequest{Name: "A", Email: "a@example.com", Password: "pass"}
    if err := svc.Register(req); err != nil {
        t.Fatalf("Register error: %v", err)
    }
    u, _ := repo.GetByEmail("a@example.com")
    if u.Role != "user" {
        t.Fatalf("expected default role 'user', got %q", u.Role)
    }
    if u.Password == "pass" || u.Password == "" {
        t.Fatalf("expected hashed password, got %q", u.Password)
    }
}

func TestUserService_Login_SuccessAndFailures(t *testing.T) {
    os.Setenv("JWT_SECRET", "secret")
    repo := &mockUserRepo{users: map[string]*domain.User{}}
    svc := NewUserService(repo)

    // register user
    _ = svc.Register(&port.RegisterRequest{Name: "U", Email: "u@mail", Password: "123"})

    // success
    resp, err := svc.Login(&port.LoginRequest{Email: "u@mail", Password: "123"})
    if err != nil || resp == nil || resp.Token == "" {
        t.Fatalf("expected token, got resp=%v err=%v", resp, err)
    }

    // wrong password
    if _, err := svc.Login(&port.LoginRequest{Email: "u@mail", Password: "wrong"}); err == nil {
        t.Fatalf("expected error on wrong password")
    }
    // unknown email
    if _, err := svc.Login(&port.LoginRequest{Email: "x@mail", Password: "123"}); err == nil {
        t.Fatalf("expected error on unknown email")
    }
}
