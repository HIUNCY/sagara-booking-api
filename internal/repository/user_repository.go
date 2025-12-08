package repository

import (
	"github.com/HIUNCY/sagara-booking-api/internal/core/domain"
	"github.com/HIUNCY/sagara-booking-api/internal/core/port"
	"gorm.io/gorm"
)

type UserRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &UserRepositoryDB{db: db}
}

func (r *UserRepositoryDB) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryDB) GetByEmail(email string) (*domain.User, error) {
	var user domain.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
