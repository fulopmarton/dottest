package mappingservice

import (
	"dottest/internal/db"
	"dottest/internal/models"
)

func AddMapping(domain string, target string) error {
	db := db.InitDB()
	err := db.Create(&models.Mapping{
		Domain: domain,
		Target: target,
	})
	return err.Error
}
