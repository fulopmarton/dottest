package store

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	if err := db.AutoMigrate(&Mapping{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
