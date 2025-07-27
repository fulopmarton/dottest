package db

import (
	"dottest/internal/models"
	"dottest/utils"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	if db != nil {
		return db
	}
	dbPath := utils.GetAppDataPath("data", "dottest.db")
	// fmt.Printf("Opening database at %s\n", dbPath)
	var err error
	db, err = gorm.Open(
		sqlite.Open(
			dbPath,
		),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	if err := db.AutoMigrate(&models.Mapping{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
