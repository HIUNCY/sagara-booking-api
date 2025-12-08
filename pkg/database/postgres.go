package database

import (
	"fmt"
	"log"
	"os"

	"github.com/HIUNCY/sagara-booking-api/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	sslMode := os.Getenv("DB_SSLMODE")
	if sslMode == "" {
		sslMode = "disable"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		sslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("✅ Database Connected to Neon Tech!")

	err = db.AutoMigrate(&domain.User{}, &domain.Field{}, &domain.Booking{})
	if err != nil {
		return nil, err
	}
	log.Println("✅ Database Migrated!")

	return db, nil
}
