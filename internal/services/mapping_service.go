package services

import (
	"dottest/internal/db"
	"dottest/internal/models"
	"log"
)

func AddMapping(domain string, target string) error {
	db := db.InitDB()
	err := db.Create(&models.Mapping{
		Domain: domain,
		Target: target,
	})
	log.Printf("%v", db)
	return err.Error
}
