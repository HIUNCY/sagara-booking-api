package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role" gorm:"default:'user'"`
}

type Field struct {
	gorm.Model
	Name         string `json:"name"`
	PricePerHour int    `json:"price_per_hour"`
	Location     string `json:"location"`
}

type Booking struct {
	gorm.Model
	FieldID   uint      `json:"field_id"`
	Field     Field     `json:"field" gorm:"foreignKey:FieldID"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status" gorm:"default:'pending'"`
}
