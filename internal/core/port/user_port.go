package port

import "github.com/HIUNCY/sagara-booking-api/internal/core/domain"

// DTO
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Repository Interface
type UserRepository interface {
	CreateUser(user *domain.User) error
	GetByEmail(email string) (*domain.User, error)
}

// Service Interface
type UserService interface {
	Register(req *RegisterRequest) error
	Login(req *LoginRequest) (*LoginResponse, error)
}
