package db

import (
	"dottest/internal/models"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(
		sqlite.Open(
			os.Getenv("DB_NAME"),
		),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	if err := db.AutoMigrate(&models.Mapping{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
