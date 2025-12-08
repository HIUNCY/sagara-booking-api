package service

import (
	"errors"

	"github.com/HIUNCY/sagara-booking-api/internal/core/domain"
	"github.com/HIUNCY/sagara-booking-api/internal/core/port"
	"github.com/HIUNCY/sagara-booking-api/pkg/util"
)

type UserServiceImpl struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) port.UserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) Register(req *port.RegisterRequest) error {
	hashedPwd, err := util.HashPassword(req.Password)
	if err != nil {
		return err
	}

	if req.Role == "" {
		req.Role = "user"
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPwd,
		Role:     req.Role,
	}

	return s.repo.CreateUser(user)
}

func (s *UserServiceImpl) Login(req *port.LoginRequest) (*port.LoginResponse, error) {
	user, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !util.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := util.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	return &port.LoginResponse{Token: token}, nil
}
